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
	svc := NewDiscordService(config.Config{
		BotVersion: "test",
		HostURL:    "https://hidereplier.example/",
		DiscordTargets: []config.DiscordTarget{
			{ID: "main", Label: "Main", WebhookURL: mainWebhook.URL, Default: true},
			{ID: "ops", Label: "Ops", WebhookURL: opsWebhook.URL},
		},
	}, fakeStore)

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
	if fakeStore.counter != 42 {
		t.Fatalf("expected counter 42, got %d", fakeStore.counter)
	}
	if fakeStore.history[0].SerialNumber != 42 {
		t.Fatalf("expected stored serial 42, got %d", fakeStore.history[0].SerialNumber)
	}
	if stored.Extras[colorKey] != "8150271" {
		t.Fatalf("expected normalized decimal color, got %q", stored.Extras[colorKey])
	}
}

func TestPostAnonymousMessageDoesNotPersistWhenWebhookFails(t *testing.T) {
	ctx := context.Background()

	webhook := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusTooManyRequests)
	}))
	defer webhook.Close()

	fakeStore := newFakeStore(10)
	svc := NewDiscordService(config.Config{
		BotVersion: "test",
		HostURL:    "https://hidereplier.example/",
		DiscordTargets: []config.DiscordTarget{
			{ID: "main", Label: "Main", WebhookURL: webhook.URL, Default: true},
		},
	}, fakeStore)

	_, err := svc.PostAnonymousMessage(ctx, model.IncomingPost{
		Content:  "hello",
		Username: "anonymous",
		Color:    "#7c5cff",
	})
	if err == nil {
		t.Fatal("expected webhook failure error")
	}
	if len(fakeStore.history) != 0 {
		t.Fatalf("expected no history rows after failed dispatch, got %d", len(fakeStore.history))
	}
	// The serial is reserved atomically before dispatch, so a failure leaves a
	// gap rather than reusing the number under a global lock.
	if fakeStore.counter != 11 {
		t.Fatalf("expected reserved serial 11, got %d", fakeStore.counter)
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
	svc := NewDiscordService(config.Config{
		BotVersion: "test",
		HostURL:    "https://hidereplier.example/",
		DiscordTargets: []config.DiscordTarget{
			{ID: "main", Label: "Main", WebhookURL: webhook.URL, Default: true},
		},
	}, fakeStore)

	_, err := svc.PostAnonymousMessage(ctx, model.IncomingPost{
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
	if fakeStore.counter != 7 {
		t.Fatalf("expected counter to remain 7, got %d", fakeStore.counter)
	}
}

type fakeStore struct {
	mu      sync.Mutex
	counter int
	history []model.StoreData
}

func newFakeStore(counter int) *fakeStore {
	return &fakeStore{counter: counter}
}

func (f *fakeStore) NextSerial(context.Context) (int, error) {
	f.mu.Lock()
	defer f.mu.Unlock()
	f.counter++
	return f.counter, nil
}

func (f *fakeStore) InsertHistory(_ context.Context, data model.StoreData) error {
	f.mu.Lock()
	defer f.mu.Unlock()
	f.history = append(f.history, data)
	return nil
}

func (f *fakeStore) StreamHistory(_ context.Context, limit int64, fn func(model.StoreData) error) error {
	f.mu.Lock()
	rows := append([]model.StoreData(nil), f.history...)
	f.mu.Unlock()

	if limit > 0 && int64(len(rows)) > limit {
		rows = rows[int64(len(rows))-limit:]
	}
	for _, row := range rows {
		if err := fn(row); err != nil {
			return err
		}
	}
	return nil
}
