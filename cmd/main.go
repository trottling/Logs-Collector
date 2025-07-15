package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"

	"log_stash_lite/internal/api"
	"log_stash_lite/internal/config"
	"log_stash_lite/internal/elastic"
	"log_stash_lite/internal/logger"
)

func main() {
	cfg := config.Load()
	log := logger.New()
	es, err := elastic.NewClient(cfg.ElasticURL, log)
	if err != nil {
		log.Fatal("failed to create elastic client", zap.Error(err))
	}

	r := chi.NewRouter()
	h := api.NewHandler(log, es)
	h.RegisterRoutes(r)

	log.Info("starting server", zap.String("addr", cfg.ListenAddr))
	if err := http.ListenAndServe(fmt.Sprintf(":%s", cfg.ListenAddr), r); err != nil {
		log.Fatal("server error", zap.Error(err))
	}
}
