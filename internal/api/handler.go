package api

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"

	"log_stash_lite/internal/elastic"
)

type Handler struct {
	log *zap.Logger
	es  *elastic.Client
}

func NewHandler(log *zap.Logger, es *elastic.Client) *Handler {
	return &Handler{log: log, es: es}
}

func (h *Handler) RegisterRoutes(r chi.Router) {
	r.Post("/add_logs", h.handleAddLogs)
	r.Get("/get_logs", h.handleGetLogs)
	r.Get("/logs_stats", h.handleLogStats)
}

func (h *Handler) handleAddLogs(w http.ResponseWriter, r *http.Request) {
	var entry map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&entry); err != nil {
		h.log.Error("invalid request", zap.Error(err))
		h.respond(w, http.StatusBadRequest, map[string]string{"error": "invalid JSON"})
		return
	}

	if err := h.es.IndexLog(entry); err != nil {
		h.log.Error("failed to index log", zap.Error(err))
		h.respond(w, http.StatusInternalServerError, map[string]string{"error": "failed to store log"})
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *Handler) handleGetLogs(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) handleLogStats(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) respond(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(data)
}
