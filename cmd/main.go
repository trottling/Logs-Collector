package main

import (
	"context"
	"log_stash_lite/internal/api/handlers"
	"log_stash_lite/internal/parser"
	"log_stash_lite/internal/storage"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"

	"log_stash_lite/internal/config"
	"log_stash_lite/internal/elastic"
	"log_stash_lite/internal/logger"
)

func main() {
	cfg := config.Load()
	log := logger.New(cfg.LogLevel)
	pr := parser.New(log)
	es, err := elastic.NewClient(cfg, log)
	if err != nil {
		log.Fatal("failed to create elastic client", zap.Error(err))
	}

	r := chi.NewRouter()
	var store storage.Storage = es
	h := handlers.NewHandler(log, store, pr)
	h.RegisterRoutes(r)

	server := &http.Server{
		Addr:    cfg.ListenAddr,
		Handler: r,
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		log.Info("starting server", zap.String("addr", cfg.ListenAddr))
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("server error", zap.Error(err))
		}
	}()

	<-quit
	log.Info("shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("server forced to shutdown", zap.Error(err))
	}

	log.Info("server exited gracefully")
}
