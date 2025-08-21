package api

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	dto "logs-collector/pkg/dto/auth"
	mw "logs-collector/pkg/middleware"
	"logs-collector/services/auth/internal/jwt"
	"logs-collector/services/auth/internal/store"
)

type Handlers struct {
	store *store.Store
	jwt   *jwt.Manager
}

func NewHandlers(s *store.Store, j *jwt.Manager) *Handlers {
	return &Handlers{store: s, jwt: j}
}

func NewRouter(st *store.Store, j *jwt.Manager) http.Handler {
	h := NewHandlers(st, j)

	r := chi.NewRouter()
	r.Use(middleware.RequestID, middleware.Recoverer, middleware.Timeout(30*time.Second))

	// health
	r.Get("/health", h.Health)

	// публичная auth-зона
	r.Route("/v1/auth", func(r chi.Router) {
		r.Post("/register", h.Register)
		r.Post("/login", h.Login)
		r.Group(func(p chi.Router) {
			p.Use(h.Authz) // JWT required
			p.Get("/me", h.Me)
		})
	})

	// защищённые CRUD-зоны
	r.Group(func(r chi.Router) {
		r.Use(h.Authz)

		// CRUD пользователей: доступно admin и root
		r.Route("/v1/users", func(r chi.Router) {
			r.With(mw.RequireRole(dto.RoleAdmin)).Post("/", h.CreateUser)
			r.With(mw.RequireRole(dto.RoleAdmin)).Get("/", h.ListUsers)
			r.With(mw.RequireRole(dto.RoleAdmin)).Get("/{id}", h.GetUser)
			r.With(mw.RequireRole(dto.RoleAdmin)).Patch("/{id}", h.UpdateUser)
			r.With(mw.RequireRole(dto.RoleAdmin)).Delete("/{id}", h.DeleteUser)
		})

		// CRUD админов: только root
		r.Route("/v1/admins", func(r chi.Router) {
			r.With(mw.RequireRole(dto.RoleRoot)).Patch("/{id}/promote", h.PromoteToAdmin) // повысить юзера до admin
			r.With(mw.RequireRole(dto.RoleRoot)).Patch("/{id}/demote", h.DemoteToUser)    // понизить админа до user
		})
	})

	return r
}
