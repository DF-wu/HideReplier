package config

import "testing"

func TestResolveDiscordTargetUsesExplicitTarget(t *testing.T) {
	cfg := Config{
		DiscordTargets: []DiscordTarget{
			{ID: "main", Label: "Main", WebhookURL: "https://example.com/main", Default: true},
			{ID: "ops", Label: "Ops", WebhookURL: "https://example.com/ops"},
		},
	}

	target, err := cfg.ResolveDiscordTarget("ops")
	if err != nil {
		t.Fatalf("resolve target: %v", err)
	}

	if target.WebhookURL != "https://example.com/ops" {
		t.Fatalf("expected ops webhook, got %s", target.WebhookURL)
	}
}

func TestResolveDiscordTargetFallsBackToDefault(t *testing.T) {
	cfg := Config{
		DiscordTargets: []DiscordTarget{
			{ID: "main", Label: "Main", WebhookURL: "https://example.com/main", Default: true},
			{ID: "ops", Label: "Ops", WebhookURL: "https://example.com/ops"},
		},
	}

	target, err := cfg.ResolveDiscordTarget("")
	if err != nil {
		t.Fatalf("resolve target: %v", err)
	}

	if target.ID != "main" {
		t.Fatalf("expected main target, got %s", target.ID)
	}
}

func TestResolveDiscordTargetRejectsUnknownTarget(t *testing.T) {
	cfg := Config{
		DiscordTargets: []DiscordTarget{
			{ID: "main", Label: "Main", WebhookURL: "https://example.com/main"},
		},
	}

	if _, err := cfg.ResolveDiscordTarget("missing"); err == nil {
		t.Fatal("expected unknown target error")
	}
}

func TestResolveDiscordTargetAcceptsDefaultForLegacyWebhook(t *testing.T) {
	cfg := Config{DiscordWebhook: "https://example.com/default"}

	target, err := cfg.ResolveDiscordTarget("default")
	if err != nil {
		t.Fatalf("resolve target: %v", err)
	}

	if target.WebhookURL != cfg.DiscordWebhook {
		t.Fatalf("expected legacy webhook, got %s", target.WebhookURL)
	}
}

func TestPublicDiscordTargetsRedactsWebhookURL(t *testing.T) {
	cfg := Config{
		DiscordTargets: []DiscordTarget{
			{ID: "main", Label: "Main", WebhookURL: "https://example.com/main", Default: true},
		},
	}

	targets := cfg.PublicDiscordTargets()
	if len(targets) != 1 {
		t.Fatalf("expected one target, got %d", len(targets))
	}

	if targets[0].WebhookURL != "" {
		t.Fatalf("expected webhook URL to be redacted, got %s", targets[0].WebhookURL)
	}
}

func TestValidateRejectsDuplicateDiscordTargetIDs(t *testing.T) {
	cfg := Config{
		DiscordTargets: []DiscordTarget{
			{ID: "main", Label: "Main", WebhookURL: "https://example.com/main"},
			{ID: "main", Label: "Duplicate", WebhookURL: "https://example.com/dup"},
		},
		MongoURI: "mongodb://localhost:27017",
	}

	if err := cfg.Validate(); err == nil {
		t.Fatal("expected duplicate target validation error")
	}
}
