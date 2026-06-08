package app

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/DF-wu/HideReplier/internal/config"
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
