package handlers

import "github.com/go-chi/chi/v5"

// RegisterRoutes registers all API routes
func (h *Handler) RegisterRoutes(r chi.Router) {
	r.Post("/add_log", h.handleAddLog)
	r.Post("/add_logs", h.handleAddLogs)
	r.Get("/get_logs", h.handleGetLogs)
	r.Get("/logs_stats", h.handleLogStats)
}
