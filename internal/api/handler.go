package api

import (
	"encoding/json"
	"net/http"

	"log_stash_lite/internal/elastic"
	"log_stash_lite/internal/parser"

	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
)

type Handler struct {
	log *zap.Logger
	pr  *parser.LogParser
	es  *elastic.Client
}

func NewHandler(log *zap.Logger, es *elastic.Client, pr *parser.LogParser) *Handler {
	return &Handler{log: log, es: es, pr: pr}
}

func (h *Handler) RegisterRoutes(r chi.Router) {
	r.Post("/add_logs", h.handleAddLogs)
	r.Get("/get_logs", h.handleGetLogs)
	r.Get("/logs_stats", h.handleLogStats)
}

func (h *Handler) handleAddLogs(w http.ResponseWriter, r *http.Request) {
	var req struct {
		ParseType string                 `json:"parse_type"`
		Log       map[string]interface{} `json:"log"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.log.Error("failed to decode request", zap.Error(err))
		h.respond(w, http.StatusBadRequest, "invalid request body")
		return
	}

	normalized, err := h.pr.Parse(req.Log, req.ParseType)
	if err != nil {
		h.log.Error("failed to parse request", zap.Error(err))
		h.respond(w, http.StatusBadRequest, "failed to parse log")
		return
	}

	if err := h.es.IndexLog(normalized); err != nil {
		h.log.Error("failed to index log", zap.Error(err))
		h.respond(w, http.StatusInternalServerError, "failed to index log")
		return
	}

	h.respond(w, http.StatusCreated, map[string]string{"status": "ok"})
}

func (h *Handler) handleGetLogs(w http.ResponseWriter, r *http.Request) {
	filters := make(map[string]string)
	for key, values := range r.URL.Query() {
		if len(values) > 0 {
			filters[key] = values[0]
		}
	}

	logs, err := h.es.GetLogs(filters)
	if err != nil {
		h.log.Error("failed to get logs", zap.Error(err))
		h.respond(w, http.StatusInternalServerError, map[string]string{"error": "failed to fetch logs"})
		return
	}

	h.respond(w, http.StatusOK, logs)
}

func (h *Handler) handleLogStats(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) respond(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(data)
}
