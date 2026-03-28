package app

import (
	"context"
	"encoding/json"
	"io"
	"io/fs"
	"net/http"
	"path"
	"strings"

	"github.com/DF-wu/HideReplier/internal/config"
	"github.com/DF-wu/HideReplier/internal/model"
	"github.com/DF-wu/HideReplier/internal/service"
	"github.com/DF-wu/HideReplier/internal/store"
)

type Handler struct {
	service    *service.DiscordService
	staticFS   http.Handler
	staticOpen fs.FS
}

// NewHandler wires the HTTP surface of the Go port around the Discord service
// and the static frontend assets.
func NewHandler(ctx context.Context, cfg config.Config, mongoStore *store.MongoStore, staticFS fs.FS) (http.Handler, error) {
	discordService, err := service.NewDiscordService(ctx, cfg, mongoStore)
	if err != nil {
		return nil, err
	}

	return &Handler{
		service:    discordService,
		staticFS:   http.FileServer(http.FS(staticFS)),
		staticOpen: staticFS,
	}, nil
}

// ServeHTTP keeps the Go port route-compatible with the existing Spring Boot app.
func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch {
	case r.URL.Path == "/actuator/health" && r.Method == http.MethodGet:
		writeJSON(w, http.StatusOK, model.HealthResponse{Status: "UP"})
	case r.URL.Path == "/HideBot/discord/version" && r.Method == http.MethodGet:
		writeJSON(w, http.StatusOK, h.service.GetVersion())
	case r.URL.Path == "/HideBot/discord" && r.Method == http.MethodGet:
		h.handleGetHistory(w, r)
	case r.URL.Path == "/HideBot/discord" && r.Method == http.MethodPost:
		h.handlePostMessage(w, r)
	case shouldServeStatic(r.URL.Path):
		h.serveStatic(w, r)
	default:
		http.NotFound(w, r)
	}
}

func (h *Handler) handleGetHistory(w http.ResponseWriter, r *http.Request) {
	history, err := h.service.GetHistory(r.Context())
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}

	writeJSON(w, http.StatusOK, history)
}

func (h *Handler) handlePostMessage(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var post model.IncomingPost
	if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid request body"})
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

func (h *Handler) serveStatic(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" || r.URL.Path == "/index.html" {
		h.serveIndexHTML(w)
		return
	}

	cleanPath := strings.TrimPrefix(path.Clean(r.URL.Path), "/")
	if _, err := fs.Stat(h.staticOpen, cleanPath); err != nil {
		http.NotFound(w, r)
		return
	}

	h.staticFS.ServeHTTP(w, r)
}

func (h *Handler) serveIndexHTML(w http.ResponseWriter) {
	file, err := h.staticOpen.Open("index.html")
	if err != nil {
		http.Error(w, "index.html not found", http.StatusNotFound)
		return
	}
	defer file.Close()

	contents, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "failed to read index.html", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(contents)
}

func shouldServeStatic(urlPath string) bool {
	if urlPath == "/" {
		return true
	}

	return strings.HasPrefix(urlPath, "/thumbs/") ||
		strings.HasSuffix(urlPath, ".css") ||
		strings.HasSuffix(urlPath, ".js") ||
		strings.HasSuffix(urlPath, ".svg") ||
		strings.HasSuffix(urlPath, ".html")
}

func writeJSON(w http.ResponseWriter, status int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(payload)
}
