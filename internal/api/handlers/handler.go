package handlers

import (
	"encoding/json"
	"log_stash_lite/internal/config"
	"net/http"

	"log_stash_lite/internal/parser"
	"log_stash_lite/internal/storage"

	"go.uber.org/zap"
)

// Handler is the main API handler
type Handler struct {
	log *zap.Logger
	pr  *parser.LogParser
	es  storage.Storage
	cfg config.Config
}

// NewHandler creates a new Handler instance
func NewHandler(log *zap.Logger, es storage.Storage, pr *parser.LogParser, cfg config.Config) *Handler {
	return &Handler{log: log, es: es, pr: pr, cfg: cfg}
}

// respond sends a JSON response
func (h *Handler) respond(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(data)
}
