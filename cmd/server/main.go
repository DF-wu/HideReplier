package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/DF-wu/HideReplier/internal/app"
	"github.com/DF-wu/HideReplier/internal/config"
	"github.com/DF-wu/HideReplier/internal/service"
	"github.com/DF-wu/HideReplier/internal/store"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	cfg := config.Load()
	if err := cfg.Validate(); err != nil {
		log.Fatalf("invalid configuration: %v", err)
	}

	mongoStore, err := store.NewMongoStore(ctx, cfg)
	if err != nil {
		log.Fatalf("failed to connect to mongodb: %v", err)
	}
	defer func() {
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		if closeErr := mongoStore.Close(shutdownCtx); closeErr != nil {
			log.Printf("failed to close mongodb client: %v", closeErr)
		}
	}()

	discordService := service.NewDiscordService(cfg, mongoStore)

	handler, err := app.NewHandler(cfg, discordService, os.DirFS(cfg.StaticDir))
	if err != nil {
		log.Fatalf("failed to initialize app (STATIC_DIR=%s): %v", cfg.StaticDir, err)
	}

	server := &http.Server{
		Addr:              cfg.ListenAddress(),
		Handler:           handler,
		ReadHeaderTimeout: 10 * time.Second,
		ReadTimeout:       30 * time.Second,
		WriteTimeout:      60 * time.Second,
		IdleTimeout:       120 * time.Second,
		MaxHeaderBytes:    16 << 10,
	}

	go func() {
		<-ctx.Done()
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		if shutdownErr := server.Shutdown(shutdownCtx); shutdownErr != nil && !errors.Is(shutdownErr, http.ErrServerClosed) {
			log.Printf("http shutdown error: %v", shutdownErr)
		}
	}()

	log.Printf("HideReplier Go server listening on %s", cfg.ListenAddress())
	if serveErr := server.ListenAndServe(); serveErr != nil && !errors.Is(serveErr, http.ErrServerClosed) {
		log.Fatalf("http server failed: %v", serveErr)
	}
}
