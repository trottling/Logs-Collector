package api

import (
	"logs-collector/pkg/dto"
	"logs-collector/pkg/dto/auth"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (h *Handlers) DeleteUser(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	actorRole := getRole(r)

	// админ может удалять только users
	if actorRole == auth_dto.RoleAdmin {
		target, err := h.store.FindByID(r.Context(), id)
		if err != nil {
			respond(w, http.StatusNotFound, dto.ErrorResponse{Error: "not found"})
			return
		}

		if target.Role != auth_dto.RoleUser {
			respond(w, http.StatusForbidden, dto.ErrorResponse{Error: "admins can delete only users"})
			return
		}
	}

	ok, err := h.store.DeleteUser(r.Context(), id)
	if err != nil {
		respond(w, http.StatusInternalServerError, dto.ErrorResponse{Error: "db error"})
		return
	}

	respond(w, http.StatusOK, auth_dto.UserDeleteResp{Deleted: ok})
}
