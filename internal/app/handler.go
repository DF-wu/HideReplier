package app

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"log"
	"mime"
	"net/http"
	"path"
	"strconv"
	"strings"
	"time"

	"github.com/DF-wu/HideReplier/internal/config"
	"github.com/DF-wu/HideReplier/internal/model"
	"github.com/DF-wu/HideReplier/internal/service"
)

const (
	// maxPostBodyBytes bounds the JSON body of POST /HideBot/discord. Discord
	// itself caps content at a few KB, so this is generous while keeping a
	// malicious client from pushing the 256MB machine into swap.
	maxPostBodyBytes = 64 << 10

	// maxHistoryLimit caps ?limit= on GET /HideBot/discord. Without ?limit=
	// the whole history is streamed, which stays memory-bounded.
	maxHistoryLimit = 5000

	cacheImmutable = "public, max-age=31536000, immutable"
	cacheStatic    = "public, max-age=86400"
	cacheNone      = "no-cache"
)

type Handler struct {
	config  config.Config
	service *service.DiscordService
	static  *staticHandler
}

// NewHandler wires the HTTP surface of the Go port around the Discord service
// and the static frontend assets.
func NewHandler(cfg config.Config, discordService *service.DiscordService, staticFS fs.FS) (http.Handler, error) {
	static, err := newStaticHandler(staticFS)
	if err != nil {
		return nil, err
	}

	return &Handler{
		config:  cfg,
		service: discordService,
		static:  static,
	}, nil
}

// ServeHTTP keeps the Go port route-compatible with the existing Spring Boot app.
func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch {
	case r.URL.Path == "/actuator/health" && r.Method == http.MethodGet:
		writeJSON(w, http.StatusOK, model.HealthResponse{Status: "UP"})
	case r.URL.Path == "/HideBot/discord/version" && r.Method == http.MethodGet:
		w.Header().Set("Cache-Control", cacheStatic)
		writeJSON(w, http.StatusOK, h.service.GetVersion())
	case r.URL.Path == "/HideBot/discord/targets" && r.Method == http.MethodGet:
		w.Header().Set("Cache-Control", cacheStatic)
		writeJSON(w, http.StatusOK, h.config.PublicDiscordTargets())
	case r.URL.Path == "/HideBot/discord" && r.Method == http.MethodGet:
		h.handleGetHistory(w, r)
	case r.URL.Path == "/HideBot/discord" && r.Method == http.MethodPost:
		h.handlePostMessage(w, r)
	case shouldServeStatic(r.URL.Path):
		h.static.ServeHTTP(w, r)
	default:
		http.NotFound(w, r)
	}
}

// handleGetHistory streams the JSON array straight from the database cursor
// so the response never materializes the whole collection in memory.
// ?limit=N returns only the N most recent rows (still in ascending order).
func (h *Handler) handleGetHistory(w http.ResponseWriter, r *http.Request) {
	limit, err := parseHistoryLimit(r.URL.Query().Get("limit"))
	if err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Cache-Control", cacheNone)

	started := false
	enc := json.NewEncoder(w)
	streamErr := h.service.StreamHistory(r.Context(), limit, func(row model.StoreData) error {
		if !started {
			started = true
			w.WriteHeader(http.StatusOK)
			if _, err := io.WriteString(w, "["); err != nil {
				return err
			}
		} else if _, err := io.WriteString(w, ","); err != nil {
			return err
		}
		return enc.Encode(row)
	})

	if streamErr != nil {
		if !started {
			writeJSON(w, http.StatusInternalServerError, map[string]string{"error": streamErr.Error()})
			return
		}
		// Headers already went out; the best we can do is cut the body short
		// so the client sees invalid JSON rather than a truncated-but-valid one.
		log.Printf("history stream aborted: %v", streamErr)
		return
	}

	if !started {
		w.WriteHeader(http.StatusOK)
		_, _ = io.WriteString(w, "[]")
		return
	}
	_, _ = io.WriteString(w, "]")
}

func parseHistoryLimit(raw string) (int64, error) {
	if raw == "" || raw == "all" {
		return 0, nil
	}

	limit, err := strconv.ParseInt(raw, 10, 64)
	if err != nil || limit < 0 {
		return 0, fmt.Errorf("limit must be a non-negative integer or \"all\"")
	}
	if limit > maxHistoryLimit {
		limit = maxHistoryLimit
	}

	return limit, nil
}

func (h *Handler) handlePostMessage(w http.ResponseWriter, r *http.Request) {
	r.Body = http.MaxBytesReader(w, r.Body, maxPostBodyBytes)
	defer r.Body.Close()

	var post model.IncomingPost
	if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
		status := http.StatusBadRequest
		var tooLarge *http.MaxBytesError
		if errors.As(err, &tooLarge) {
			status = http.StatusRequestEntityTooLarge
		}
		writeJSON(w, status, map[string]string{"error": "invalid request body"})
		return
	}

	if post.Username == "" || post.Content == "" {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "username and content are required"})
		return
	}

	result, err := h.service.PostAnonymousMessage(r.Context(), post)
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}

	writeJSON(w, http.StatusOK, result)
}

func shouldServeStatic(urlPath string) bool {
	if urlPath == "/" {
		return true
	}

	return strings.HasPrefix(urlPath, "/thumbs/") ||
		strings.HasPrefix(urlPath, "/assets/") ||
		strings.HasSuffix(urlPath, ".css") ||
		strings.HasSuffix(urlPath, ".js") ||
		strings.HasSuffix(urlPath, ".gif") ||
		strings.HasSuffix(urlPath, ".ico") ||
		strings.HasSuffix(urlPath, ".png") ||
		strings.HasSuffix(urlPath, ".svg") ||
		strings.HasSuffix(urlPath, ".webp") ||
		strings.HasSuffix(urlPath, ".html")
}

func writeJSON(w http.ResponseWriter, status int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(payload)
}

// staticHandler serves the Vite build output and thumbnail assets.
//
//   - index.html is read once at boot and served from memory.
//   - Hashed Vite assets under /assets/ get a one-year immutable cache so
//     returning visitors never re-download the bundle.
//   - When the build produced a sibling ".gz" file (see Dockerfile) and the
//     client accepts gzip, the precompressed file is sent as-is. Nothing is
//     compressed at request time, which matters on a single shared vCPU.
type staticHandler struct {
	fsys      fs.FS
	index     []byte
	indexTime time.Time
}

func newStaticHandler(fsys fs.FS) (*staticHandler, error) {
	index, err := fs.ReadFile(fsys, "index.html")
	if err != nil {
		return nil, fmt.Errorf("read index.html: %w", err)
	}

	return &staticHandler{fsys: fsys, index: index, indexTime: time.Now()}, nil
}

func (s *staticHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet && r.Method != http.MethodHead {
		w.Header().Set("Allow", "GET, HEAD")
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if r.URL.Path == "/" || r.URL.Path == "/index.html" {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.Header().Set("Cache-Control", cacheNone)
		http.ServeContent(w, r, "index.html", s.indexTime, bytes.NewReader(s.index))
		return
	}

	cleanPath := strings.TrimPrefix(path.Clean(r.URL.Path), "/")
	if cleanPath == "" || cleanPath == "." || strings.HasSuffix(cleanPath, ".gz") {
		http.NotFound(w, r)
		return
	}

	info, err := fs.Stat(s.fsys, cleanPath)
	if err != nil || info.IsDir() {
		http.NotFound(w, r)
		return
	}

	contentType := contentTypeFor(cleanPath)
	w.Header().Set("Content-Type", contentType)
	if strings.HasPrefix(cleanPath, "assets/") {
		w.Header().Set("Cache-Control", cacheImmutable)
	} else {
		w.Header().Set("Cache-Control", cacheStatic)
	}

	servePath := cleanPath
	if acceptsGzip(r) {
		if gzInfo, gzErr := fs.Stat(s.fsys, cleanPath+".gz"); gzErr == nil && !gzInfo.IsDir() {
			servePath = cleanPath + ".gz"
			info = gzInfo
			w.Header().Set("Content-Encoding", "gzip")
			w.Header().Add("Vary", "Accept-Encoding")
		}
	}

	file, err := s.fsys.Open(servePath)
	if err != nil {
		http.NotFound(w, r)
		return
	}
	defer file.Close()

	if seeker, ok := file.(io.ReadSeeker); ok {
		http.ServeContent(w, r, cleanPath, info.ModTime(), seeker)
		return
	}

	w.Header().Set("Content-Length", strconv.FormatInt(info.Size(), 10))
	w.WriteHeader(http.StatusOK)
	if r.Method != http.MethodHead {
		_, _ = io.Copy(w, file)
	}
}

// extraContentTypes covers extensions missing from Go's builtin mime table
// (Alpine ships no /etc/mime.types).
var extraContentTypes = map[string]string{
	".ico":   "image/x-icon",
	".woff":  "font/woff",
	".woff2": "font/woff2",
	".txt":   "text/plain; charset=utf-8",
}

func contentTypeFor(name string) string {
	ext := strings.ToLower(path.Ext(name))
	if ct := mime.TypeByExtension(ext); ct != "" {
		return ct
	}
	if ct, ok := extraContentTypes[ext]; ok {
		return ct
	}
	return "application/octet-stream"
}

func acceptsGzip(r *http.Request) bool {
	for _, part := range strings.Split(r.Header.Get("Accept-Encoding"), ",") {
		enc := strings.TrimSpace(part)
		if i := strings.IndexByte(enc, ';'); i >= 0 {
			enc = strings.TrimSpace(enc[:i])
		}
		if enc == "gzip" || enc == "*" {
			return true
		}
	}
	return false
}
