package api

import (
	"logs-collector/pkg/dto"
	"logs-collector/pkg/dto/auth"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (h *Handlers) GetUser(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	u, err := h.store.FindByID(r.Context(), id)
	if err != nil {
		respond(w, 404, dto.ErrorResponse{Error: "not found"})
		return
	}
	respond(w, 200, auth_dto.UserResp{ID: u.ID, Login: u.Login, Role: u.Role})
}
