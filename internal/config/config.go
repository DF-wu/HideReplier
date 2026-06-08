package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type DiscordTarget struct {
	ID         string `json:"id"`
	Label      string `json:"label"`
	WebhookURL string `json:"webhookUrl"`
	Default    bool   `json:"default,omitempty"`
}

type Config struct {
	Port                     string
	BotVersion               string
	HostURL                  string
	DiscordWebhook           string
	DiscordTargets           []DiscordTarget
	DiscordTargetsParseError error
	MongoURI                 string
	MongoDatabase            string
	StaticDir                string
}

// Load reads runtime configuration from environment variables and applies
// safe local defaults only for non-secret values.
func Load() Config {
	discordTargets, discordTargetsParseError := loadDiscordTargets()

	return Config{
		Port:                     getenv("PORT", "8082"),
		BotVersion:               getenv("BOT_VERSION", "dev"),
		HostURL:                  getenv("HOST_URL", "https://hidereplier.fly.dev/"),
		DiscordWebhook:           os.Getenv("DC_WEBHOOK_URL"),
		DiscordTargets:           discordTargets,
		DiscordTargetsParseError: discordTargetsParseError,
		MongoURI:                 os.Getenv("MONGO_URI"),
		MongoDatabase:            getenv("MONGO_DATABASE", "DiscordBotHistoryDB"),
		StaticDir:                getenv("STATIC_DIR", "src/main/resources/static"),
	}
}

// Validate ensures the Go port fails fast when required runtime secrets are missing.
func (c Config) Validate() error {
	if c.DiscordTargetsParseError != nil {
		return fmt.Errorf("DISCORD_TARGETS is invalid JSON: %w", c.DiscordTargetsParseError)
	}

	if c.DiscordWebhook == "" && len(c.DiscordTargets) == 0 {
		return fmt.Errorf("DC_WEBHOOK_URL or DISCORD_TARGETS is required")
	}

	seen := make(map[string]struct{}, len(c.DiscordTargets))
	defaultCount := 0
	for _, target := range c.DiscordTargets {
		if target.ID == "" {
			return fmt.Errorf("DISCORD_TARGETS contains a target without id")
		}
		if target.Label == "" {
			return fmt.Errorf("DISCORD_TARGETS target %q is missing label", target.ID)
		}
		if target.WebhookURL == "" {
			return fmt.Errorf("DISCORD_TARGETS target %q is missing webhookUrl", target.ID)
		}
		if _, ok := seen[target.ID]; ok {
			return fmt.Errorf("DISCORD_TARGETS contains duplicate id %q", target.ID)
		}
		seen[target.ID] = struct{}{}
		if target.Default {
			defaultCount++
		}
	}

	if defaultCount > 1 {
		return fmt.Errorf("DISCORD_TARGETS can contain at most one default target")
	}

	if c.MongoURI == "" {
		return fmt.Errorf("MONGO_URI is required")
	}

	return nil
}

// ListenAddress returns the HTTP bind address used by the server.
func (c Config) ListenAddress() string {
	return fmt.Sprintf("0.0.0.0:%s", c.Port)
}

func (c Config) ResolveDiscordTarget(id string) (DiscordTarget, error) {
	if id != "" {
		for _, target := range c.DiscordTargets {
			if target.ID == id {
				return target, nil
			}
		}

		if id == "default" && c.DiscordWebhook != "" {
			return DiscordTarget{
				ID:         "default",
				Label:      "Default Discord channel",
				WebhookURL: c.DiscordWebhook,
				Default:    true,
			}, nil
		}

		return DiscordTarget{}, fmt.Errorf("unknown discord target %q", id)
	}

	for _, target := range c.DiscordTargets {
		if target.Default {
			return target, nil
		}
	}

	if len(c.DiscordTargets) > 0 {
		return c.DiscordTargets[0], nil
	}

	return DiscordTarget{
		ID:         "default",
		Label:      "Default Discord channel",
		WebhookURL: c.DiscordWebhook,
		Default:    true,
	}, nil
}

func (c Config) PublicDiscordTargets() []DiscordTarget {
	if len(c.DiscordTargets) > 0 {
		targets := make([]DiscordTarget, len(c.DiscordTargets))
		for i, target := range c.DiscordTargets {
			target.WebhookURL = ""
			targets[i] = target
		}
		return targets
	}

	if c.DiscordWebhook == "" {
		return nil
	}

	return []DiscordTarget{
		{
			ID:      "default",
			Label:   "Default Discord channel",
			Default: true,
		},
	}
}

func loadDiscordTargets() ([]DiscordTarget, error) {
	raw := os.Getenv("DISCORD_TARGETS")
	if raw == "" {
		return nil, nil
	}

	var targets []DiscordTarget
	if err := json.Unmarshal([]byte(raw), &targets); err != nil {
		return nil, err
	}

	return targets, nil
}

func getenv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}

	return fallback
}
