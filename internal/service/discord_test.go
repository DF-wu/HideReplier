package service

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"

	"github.com/DF-wu/HideReplier/internal/config"
	"github.com/DF-wu/HideReplier/internal/model"
)

func TestNormalizeColor(t *testing.T) {
	got, err := normalizeColor("#7c5cff")
	if err != nil {
		t.Fatalf("normalizeColor returned error: %v", err)
	}

	if got != 8150271 {
		t.Fatalf("expected 8150271, got %d", got)
	}
}

func TestNormalizeColorRejectsInvalidHex(t *testing.T) {
	if _, err := normalizeColor("#not-a-color"); err == nil {
		t.Fatal("expected error for invalid color")
	}
}

func TestPostAnonymousMessageDispatchesSelectedTargetAndRedactsWebhook(t *testing.T) {
	ctx := context.Background()
	targetCalls := make(map[string]int)
	var mu sync.Mutex

	mainWebhook := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		mu.Lock()
		targetCalls["main"]++
		mu.Unlock()
		w.WriteHeader(http.StatusNoContent)
	}))
	defer mainWebhook.Close()

	opsWebhook := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()

		var payload model.DiscordWebhookPayload
		if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
			t.Errorf("decode webhook payload: %v", err)
		}

		if payload.Username != "anonymous" {
			t.Errorf("expected webhook username anonymous, got %q", payload.Username)
		}
		if len(payload.Embeds) != 1 || payload.Embeds[0].Description != "hello from ops" {
			t.Errorf("unexpected embed payload: %+v", payload.Embeds)
		}

		mu.Lock()
		targetCalls["ops"]++
		mu.Unlock()
		w.WriteHeader(http.StatusNoContent)
	}))
	defer opsWebhook.Close()

	fakeStore := newFakeStore(41)
	svc, err := NewDiscordService(ctx, config.Config{
		BotVersion: "test",
		HostURL:    "https://hidereplier.example/",
		DiscordTargets: []config.DiscordTarget{
			{ID: "main", Label: "Main", WebhookURL: mainWebhook.URL, Default: true},
			{ID: "ops", Label: "Ops", WebhookURL: opsWebhook.URL},
		},
	}, fakeStore)
	if err != nil {
		t.Fatalf("new service: %v", err)
	}

	result, err := svc.PostAnonymousMessage(ctx, model.IncomingPost{
		Content:   "hello from ops",
		Username:  "anonymous",
		Color:     "#7c5cff",
		IP:        "127.0.0.1",
		Thumbnail: "https://example.com/thumb.gif",
		TargetID:  "ops",
	})
	if err != nil {
		t.Fatalf("post anonymous message: %v", err)
	}

	mu.Lock()
	mainCalls := targetCalls["main"]
	opsCalls := targetCalls["ops"]
	mu.Unlock()

	if mainCalls != 0 {
		t.Fatalf("expected main webhook to be untouched, got %d calls", mainCalls)
	}
	if opsCalls != 1 {
		t.Fatalf("expected ops webhook once, got %d calls", opsCalls)
	}
	if result.TargetID != "ops" {
		t.Fatalf("expected result target ops, got %q", result.TargetID)
	}
	if result.URL != "https://hidereplier.example/" {
		t.Fatalf("expected public host URL, got %q", result.URL)
	}
	if result.URL == opsWebhook.URL {
		t.Fatal("webhook URL leaked into API result")
	}

	if len(fakeStore.history) != 1 {
		t.Fatalf("expected one history row, got %d", len(fakeStore.history))
	}
	stored := fakeStore.history[0].DiscordMessage
	if stored.URL == opsWebhook.URL {
		t.Fatal("webhook URL leaked into stored history")
	}
	if stored.TargetID != "ops" {
		t.Fatalf("expected stored target ops, got %q", stored.TargetID)
	}
	if fakeStore.counter.Counter != 42 {
		t.Fatalf("expected counter 42, got %d", fakeStore.counter.Counter)
	}
}

func TestPostAnonymousMessageRejectsUnknownTargetBeforeDispatchAndPersistence(t *testing.T) {
	ctx := context.Background()
	dispatchCalled := false

	webhook := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		dispatchCalled = true
		w.WriteHeader(http.StatusNoContent)
	}))
	defer webhook.Close()

	fakeStore := newFakeStore(7)
	svc, err := NewDiscordService(ctx, config.Config{
		BotVersion: "test",
		HostURL:    "https://hidereplier.example/",
		DiscordTargets: []config.DiscordTarget{
			{ID: "main", Label: "Main", WebhookURL: webhook.URL, Default: true},
		},
	}, fakeStore)
	if err != nil {
		t.Fatalf("new service: %v", err)
	}

	_, err = svc.PostAnonymousMessage(ctx, model.IncomingPost{
		Content:  "hello",
		Username: "anonymous",
		Color:    "#7c5cff",
		TargetID: "missing",
	})
	if err == nil {
		t.Fatal("expected unknown target error")
	}

	if dispatchCalled {
		t.Fatal("webhook dispatch should not run for unknown target")
	}
	if len(fakeStore.history) != 0 {
		t.Fatalf("expected no history rows, got %d", len(fakeStore.history))
	}
	if fakeStore.counter.Counter != 7 {
		t.Fatalf("expected counter to remain 7, got %d", fakeStore.counter.Counter)
	}
}

type fakeStore struct {
	counter *model.SerialCounter
	history []model.StoreData
}

func newFakeStore(counter int) *fakeStore {
	return &fakeStore{counter: &model.SerialCounter{ID: "counter", Counter: counter}}
}

func (f *fakeStore) LoadOrCreateCounter(context.Context) (*model.SerialCounter, error) {
	return f.counter, nil
}

func (f *fakeStore) SaveCounter(_ context.Context, counter *model.SerialCounter) error {
	f.counter = counter
	return nil
}

func (f *fakeStore) InsertHistory(_ context.Context, data model.StoreData) error {
	f.history = append(f.history, data)
	return nil
}

func (f *fakeStore) ListHistory(context.Context) ([]model.StoreData, error) {
	return f.history, nil
}
