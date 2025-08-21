package api

import (
	"logs-collector/pkg/dto"
	"logs-collector/pkg/dto/auth"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (h *Handlers) PromoteToAdmin(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	ok, err := h.store.SetRole(r.Context(), id, auth_dto.RoleAdmin)
	if err != nil {
		respond(w, 500, dto.ErrorResponse{Error: "db error"})
		return
	}
	respond(w, 200, dto.OkResp{Ok: ok})
}
