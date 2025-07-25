package main

import (
	"context"
	"errors"
	"logs-collector/internal/api/handlers"
	"logs-collector/internal/parser"
	"logs-collector/internal/storage"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "logs-collector/docs"
	"logs-collector/internal/config"
	"logs-collector/internal/elastic"
	"logs-collector/internal/logger"

	"go.uber.org/zap"
)

func main() {
	cfg := config.Load()
	log := logger.New(cfg.LogLevel)
	pr := parser.New(log)
	es, err := elastic.NewClient(cfg, log)
	if err != nil {
		log.Fatal("failed to create elastic client", zap.Error(err))
	}

	var store storage.Storage = es
	h := handlers.NewHandler(log, store, pr, cfg)

	jwtSecret := cfg.JWTSecret
	if jwtSecret == "" {
		log.Fatal("JWT_SECRET env variable is required")
	}

	r := handlers.NewRouter(h, []byte(jwtSecret))

	server := &http.Server{
		Addr:         cfg.ListenAddr,
		Handler:      r,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		log.Info("starting server", zap.String("addr", cfg.ListenAddr))
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
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
