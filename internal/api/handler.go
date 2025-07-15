package api

import (
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
	r.Post("/logs", h.handlePostLog)
	r.Get("/logs", h.handleGetLogs)
}

func (h *Handler) handlePostLog(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
}

func (h *Handler) handleGetLogs(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
