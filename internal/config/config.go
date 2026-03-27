package config

import (
	"fmt"
	"os"
)

type Config struct {
	Port           string
	BotVersion     string
	HostURL        string
	DiscordWebhook string
	MongoURI       string
	MongoDatabase  string
	StaticDir      string
}

func Load() Config {
	return Config{
		Port:           getenv("PORT", "8082"),
		BotVersion:     getenv("BOT_VERSION", "dev"),
		HostURL:        getenv("HOST_URL", "https://hidereplier.fly.dev/"),
		DiscordWebhook: os.Getenv("DC_WEBHOOK_URL"),
		MongoURI:       os.Getenv("MONGO_URI"),
		MongoDatabase:  getenv("MONGO_DATABASE", "DiscordBotHistoryDB"),
		StaticDir:      getenv("STATIC_DIR", "src/main/resources/static"),
	}
}

func (c Config) Validate() error {
	if c.DiscordWebhook == "" {
		return fmt.Errorf("DC_WEBHOOK_URL is required")
	}

	if c.MongoURI == "" {
		return fmt.Errorf("MONGO_URI is required")
	}

	return nil
}

func (c Config) ListenAddress() string {
	return fmt.Sprintf("0.0.0.0:%s", c.Port)
}

func getenv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}

	return fallback
}
