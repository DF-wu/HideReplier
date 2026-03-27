package main

import (
	"context"
	"errors"
	"io/fs"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/DF-wu/HideReplier/internal/app"
	"github.com/DF-wu/HideReplier/internal/config"
	"github.com/DF-wu/HideReplier/internal/store"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	cfg := config.Load()
	if err := cfg.Validate(); err != nil {
		log.Fatalf("invalid configuration: %v", err)
	}

	staticFS := os.DirFS(cfg.StaticDir)
	if _, err := fs.Stat(staticFS, "index.html"); err != nil {
		log.Fatalf("failed to locate static frontend assets in %s: %v", cfg.StaticDir, err)
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

	handler, err := app.NewHandler(ctx, cfg, mongoStore, staticFS)
	if err != nil {
		log.Fatalf("failed to initialize app: %v", err)
	}

	server := &http.Server{
		Addr:              cfg.ListenAddress(),
		Handler:           handler,
		ReadHeaderTimeout: 10 * time.Second,
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
