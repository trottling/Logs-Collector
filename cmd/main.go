package main

import (
	"fmt"
	"log_stash_lite/internal/api/handlers"
	"log_stash_lite/internal/parser"
	"net/http"

	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"

	"log_stash_lite/internal/config"
	"log_stash_lite/internal/elastic"
	"log_stash_lite/internal/logger"
)

func main() {
	cfg := config.Load()
	log := logger.New()
	pr := parser.New(log)
	es, err := elastic.NewClient(cfg.ElasticURL, log)
	if err != nil {
		log.Fatal("failed to create elastic client", zap.Error(err))
	}

	r := chi.NewRouter()
	h := handlers.NewHandler(log, es, pr)
	h.RegisterRoutes(r)

	log.Info("starting server", zap.String("addr", cfg.ListenAddr))
	if err := http.ListenAndServe(fmt.Sprintf(":%s", cfg.ListenAddr), r); err != nil {
		log.Fatal("server error", zap.Error(err))
	}
}
