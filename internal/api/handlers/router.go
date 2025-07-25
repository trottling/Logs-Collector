package handlers

import (
	"log_stash_lite/internal/api/middleware"

	"github.com/go-chi/chi/v5"
	httpSwagger "github.com/swaggo/http-swagger"
)

// NewRouter creates and configures a router with public and protected routes
func NewRouter(h *Handler, jwtSecret []byte) *chi.Mux {
	r := chi.NewRouter()

	// Public routes
	r.Get("/swagger/*", httpSwagger.WrapHandler)
	r.Get("/auth/token", h.HandleAuthToken)
	r.Get("/health", h.HandleHealth)

	// Protected routes
	r.Group(func(r chi.Router) {
		r.Use(middleware.JWTAuthMiddleware(jwtSecret))
		r.Post("/add_log", h.HandleAddLog)
		r.Post("/add_logs", h.HandleAddLogs)
		r.Get("/get_logs", h.HandleGetLogs)
		r.Get("/get_logs_count", h.HandleGetLogsCount)
	})

	return r
}
