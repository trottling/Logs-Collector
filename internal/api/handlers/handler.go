package handlers

import (
	"encoding/json"
	"net/http"

	"log_stash_lite/internal/elastic"
	"log_stash_lite/internal/parser"

	"go.uber.org/zap"
)

var okResp = map[string]string{"status": "ok"}

type Handler struct {
	log *zap.Logger
	pr  *parser.LogParser
	es  *elastic.Client
}

func NewHandler(log *zap.Logger, es *elastic.Client, pr *parser.LogParser) *Handler {
	return &Handler{log: log, es: es, pr: pr}
}

func (h *Handler) handleLogStats(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) respond(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(data)
}
