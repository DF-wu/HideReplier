package app

import (
	"bytes"
	"compress/gzip"
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"testing/fstest"

	"github.com/DF-wu/HideReplier/internal/config"
	"github.com/DF-wu/HideReplier/internal/model"
	"github.com/DF-wu/HideReplier/internal/service"
)

func TestShouldServeViteStaticAssets(t *testing.T) {
	paths := []string{
		"/",
		"/assets/index-bc5c4807.js",
		"/assets/index-2bffa169.css",
		"/icon.png",
		"/thumbs/01.gif",
	}

	for _, path := range paths {
		t.Run(path, func(t *testing.T) {
			if !shouldServeStatic(path) {
				t.Fatalf("expected %s to be served as a static asset", path)
			}
		})
	}
}

func TestShouldNotServeUnknownStaticPath(t *testing.T) {
	if shouldServeStatic("/HideBot/not-a-static-route") {
		t.Fatal("expected API-looking paths to stay out of static routing")
	}
}

func TestDiscordTargetsRouteRedactsWebhookURLs(t *testing.T) {
	handler := &Handler{
		config: config.Config{
			DiscordTargets: []config.DiscordTarget{
				{
					ID:         "main",
					Label:      "Main",
					WebhookURL: "https://discord.example/webhook",
					Default:    true,
				},
			},
		},
	}

	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, "/HideBot/discord/targets", nil)
	handler.ServeHTTP(recorder, request)

	if recorder.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", recorder.Code)
	}

	var targets []config.DiscordTarget
	if err := json.NewDecoder(recorder.Body).Decode(&targets); err != nil {
		t.Fatalf("decode targets: %v", err)
	}

	if len(targets) != 1 {
		t.Fatalf("expected one target, got %d", len(targets))
	}

	if targets[0].WebhookURL != "" {
		t.Fatalf("expected webhook URL to be redacted, got %s", targets[0].WebhookURL)
	}
}

func gzipBytes(t *testing.T, raw string) []byte {
	t.Helper()
	var buf bytes.Buffer
	zw := gzip.NewWriter(&buf)
	if _, err := zw.Write([]byte(raw)); err != nil {
		t.Fatal(err)
	}
	if err := zw.Close(); err != nil {
		t.Fatal(err)
	}
	return buf.Bytes()
}

func newTestStatic(t *testing.T) *staticHandler {
	t.Helper()
	const bundle = "console.log('bundle');"
	fsys := fstest.MapFS{
		"index.html":                {Data: []byte("<html>index</html>")},
		"assets/index-abc123.js":    {Data: []byte(bundle)},
		"assets/index-abc123.js.gz": {Data: gzipBytes(t, bundle)},
		"thumbs/01.svg":             {Data: []byte("<svg/>")},
		"icon.ico":                  {Data: []byte("ico")},
	}
	static, err := newStaticHandler(fsys)
	if err != nil {
		t.Fatalf("new static handler: %v", err)
	}
	return static
}

func TestStaticIndexServedFromMemoryWithNoCache(t *testing.T) {
	handler := &Handler{static: newTestStatic(t)}

	for _, p := range []string{"/", "/index.html"} {
		rec := httptest.NewRecorder()
		handler.ServeHTTP(rec, httptest.NewRequest(http.MethodGet, p, nil))
		if rec.Code != http.StatusOK {
			t.Fatalf("%s: expected 200, got %d", p, rec.Code)
		}
		if got := rec.Body.String(); got != "<html>index</html>" {
			t.Fatalf("%s: unexpected body %q", p, got)
		}
		if got := rec.Header().Get("Cache-Control"); got != cacheNone {
			t.Fatalf("%s: expected Cache-Control %q, got %q", p, cacheNone, got)
		}
		if ct := rec.Header().Get("Content-Type"); !strings.HasPrefix(ct, "text/html") {
			t.Fatalf("%s: unexpected content type %q", p, ct)
		}
	}
}

func TestStaticHashedAssetsAreImmutableAndPrecompressed(t *testing.T) {
	handler := &Handler{static: newTestStatic(t)}

	req := httptest.NewRequest(http.MethodGet, "/assets/index-abc123.js", nil)
	req.Header.Set("Accept-Encoding", "br, gzip;q=0.9")
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", rec.Code)
	}
	if got := rec.Header().Get("Content-Encoding"); got != "gzip" {
		t.Fatalf("expected gzip encoding, got %q", got)
	}
	if got := rec.Header().Get("Vary"); got != "Accept-Encoding" {
		t.Fatalf("expected Vary Accept-Encoding, got %q", got)
	}
	if got := rec.Header().Get("Cache-Control"); got != cacheImmutable {
		t.Fatalf("expected immutable cache, got %q", got)
	}
	if ct := rec.Header().Get("Content-Type"); !strings.Contains(ct, "javascript") {
		t.Fatalf("expected javascript content type, got %q", ct)
	}
	zr, err := gzip.NewReader(rec.Body)
	if err != nil {
		t.Fatalf("body is not gzip: %v", err)
	}
	decoded, _ := io.ReadAll(zr)
	if string(decoded) != "console.log('bundle');" {
		t.Fatalf("unexpected decoded body %q", decoded)
	}

	// Clients that do not accept gzip get the plain file.
	rec = httptest.NewRecorder()
	handler.ServeHTTP(rec, httptest.NewRequest(http.MethodGet, "/assets/index-abc123.js", nil))
	if rec.Header().Get("Content-Encoding") != "" {
		t.Fatal("expected identity encoding without Accept-Encoding")
	}
	if rec.Body.String() != "console.log('bundle');" {
		t.Fatalf("unexpected plain body %q", rec.Body.String())
	}
}

func TestStaticHidesGzipSiblingsAndMissingFiles(t *testing.T) {
	handler := &Handler{static: newTestStatic(t)}

	for _, p := range []string{"/assets/index-abc123.js.gz", "/assets/missing.js", "/thumbs/../index.html.gz"} {
		rec := httptest.NewRecorder()
		handler.ServeHTTP(rec, httptest.NewRequest(http.MethodGet, p, nil))
		if rec.Code != http.StatusNotFound {
			t.Fatalf("%s: expected 404, got %d", p, rec.Code)
		}
	}

	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, httptest.NewRequest(http.MethodGet, "/icon.ico", nil))
	// The host may have /etc/mime.types with its own ico mapping; either way
	// it must not fall back to octet-stream.
	if rec.Code != http.StatusOK || !strings.HasPrefix(rec.Header().Get("Content-Type"), "image/") {
		t.Fatalf("expected ico content type, got %d %q", rec.Code, rec.Header().Get("Content-Type"))
	}
	if rec.Header().Get("Cache-Control") != cacheStatic {
		t.Fatalf("expected static cache for non-hashed asset, got %q", rec.Header().Get("Cache-Control"))
	}
}

type historyStore struct {
	rows []model.StoreData
	err  error
}

func (h *historyStore) NextSerial(context.Context) (int, error) { return 0, nil }
func (h *historyStore) InsertHistory(context.Context, model.StoreData) error {
	return nil
}
func (h *historyStore) StreamHistory(_ context.Context, limit int64, fn func(model.StoreData) error) error {
	rows := h.rows
	if limit > 0 && int64(len(rows)) > limit {
		rows = rows[int64(len(rows))-limit:]
	}
	for _, row := range rows {
		if err := fn(row); err != nil {
			return err
		}
	}
	return h.err
}

func newHistoryHandler(store service.Store) *Handler {
	return &Handler{service: service.NewDiscordService(config.Config{}, store)}
}

func TestHistoryStreamsValidJSONArray(t *testing.T) {
	handler := newHistoryHandler(&historyStore{rows: []model.StoreData{
		{SerialNumber: 1, PosterIP: "a"},
		{SerialNumber: 2, PosterIP: "b"},
		{SerialNumber: 3, PosterIP: "c"},
	}})

	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, httptest.NewRequest(http.MethodGet, "/HideBot/discord", nil))
	if rec.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", rec.Code)
	}
	var rows []model.StoreData
	if err := json.NewDecoder(rec.Body).Decode(&rows); err != nil {
		t.Fatalf("response is not a JSON array: %v\n%s", err, rec.Body.String())
	}
	if len(rows) != 3 || rows[0].SerialNumber != 1 || rows[2].SerialNumber != 3 {
		t.Fatalf("unexpected rows %+v", rows)
	}

	rec = httptest.NewRecorder()
	handler.ServeHTTP(rec, httptest.NewRequest(http.MethodGet, "/HideBot/discord?limit=2", nil))
	rows = nil
	if err := json.NewDecoder(rec.Body).Decode(&rows); err != nil {
		t.Fatalf("limited response is not a JSON array: %v", err)
	}
	if len(rows) != 2 || rows[0].SerialNumber != 2 {
		t.Fatalf("expected the two most recent rows, got %+v", rows)
	}
}

func TestHistoryEmptyAndErrorCases(t *testing.T) {
	rec := httptest.NewRecorder()
	newHistoryHandler(&historyStore{}).ServeHTTP(rec, httptest.NewRequest(http.MethodGet, "/HideBot/discord", nil))
	if rec.Code != http.StatusOK || strings.TrimSpace(rec.Body.String()) != "[]" {
		t.Fatalf("expected empty array, got %d %q", rec.Code, rec.Body.String())
	}

	rec = httptest.NewRecorder()
	newHistoryHandler(&historyStore{err: errors.New("boom")}).ServeHTTP(rec, httptest.NewRequest(http.MethodGet, "/HideBot/discord", nil))
	if rec.Code != http.StatusInternalServerError {
		t.Fatalf("expected 500 when the cursor fails before any row, got %d", rec.Code)
	}

	rec = httptest.NewRecorder()
	newHistoryHandler(&historyStore{}).ServeHTTP(rec, httptest.NewRequest(http.MethodGet, "/HideBot/discord?limit=-1", nil))
	if rec.Code != http.StatusBadRequest {
		t.Fatalf("expected 400 for negative limit, got %d", rec.Code)
	}
}

func TestPostRejectsOversizedBody(t *testing.T) {
	handler := newHistoryHandler(&historyStore{})
	body := strings.NewReader(`{"username":"a","content":"` + strings.Repeat("x", maxPostBodyBytes+1) + `"}`)
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, httptest.NewRequest(http.MethodPost, "/HideBot/discord", body))
	if rec.Code != http.StatusRequestEntityTooLarge {
		t.Fatalf("expected 413, got %d", rec.Code)
	}
}
