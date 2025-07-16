package handlers

import (
	"encoding/json"
	"net/http"

	"log_stash_lite/internal/elastic"
	"log_stash_lite/internal/parser"

	"go.uber.org/zap"
)

// Handler is the main API handler
type Handler struct {
	log *zap.Logger
	pr  *parser.LogParser
	es  *elastic.Client
}

// NewHandler creates a new Handler instance
func NewHandler(log *zap.Logger, es *elastic.Client, pr *parser.LogParser) *Handler {
	return &Handler{log: log, es: es, pr: pr}
}

// handleLogStats returns log statistics (stub)
func (h *Handler) handleLogStats(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

// respond sends a JSON response
func (h *Handler) respond(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(data)
}
