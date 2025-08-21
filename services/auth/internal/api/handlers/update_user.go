package api

import (
	"encoding/json"
	"logs-collector/pkg/dto"
	"logs-collector/pkg/dto/auth"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (h *Handlers) UpdateUser(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	var req auth_dto.UserUpdateReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respond(w, 400, dto.ErrorResponse{Error: "invalid json"})
		return
	}

	actorRole := getRole(r)
	// admin не может менять роль на admin/root и трогать админов/рутов
	if actorRole == auth_dto.RoleAdmin {
		if req.Role != nil && *req.Role != auth_dto.RoleUser {
			respond(w, 403, dto.ErrorResponse{Error: "admins cannot set role != user"})
			return
		}
		// запрет админу редактировать не-user
		target, err := h.store.FindByID(r.Context(), id)
		if err != nil {
			respond(w, 404, dto.ErrorResponse{Error: "not found"})
			return
		}
		if target.Role != auth_dto.RoleUser {
			respond(w, 403, dto.ErrorResponse{Error: "admins can modify only users"})
			return
		}
	}

	updated, err := h.store.UpdateUser(r.Context(), id, req)
	if err != nil {
		respond(w, 500, dto.ErrorResponse{Error: "db error"})
		return
	}
	respond(w, 200, auth_dto.UserUpdateResp{Updated: updated})
}
