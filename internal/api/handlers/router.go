package handlers

import "github.com/go-chi/chi/v5"

func (h *Handler) RegisterRoutes(r chi.Router) {
	r.Post("/add_log", h.handleAddLog)
	r.Get("/get_logs", h.handleGetLogs)
	r.Get("/logs_stats", h.handleLogStats)
}
