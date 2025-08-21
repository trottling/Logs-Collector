package api

import (
	"encoding/json"
	"logs-collector/pkg/dto"
	"logs-collector/pkg/dto/auth"
	"net/http"
)

func (h *Handlers) CreateUser(w http.ResponseWriter, r *http.Request) {
	var req auth_dto.UserCreateReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respond(w, http.StatusBadRequest, dto.ErrorResponse{Error: "invalid json"})
		return
	}

	// admin может создавать только role=user
	role := getRole(r)
	if role == auth_dto.RoleAdmin && req.Role != auth_dto.RoleUser {
		respond(w, http.StatusForbidden, dto.ErrorResponse{Error: "admins can only create users"})
		return
	}

	_, err := h.store.CreateUser(r.Context(), req.Login, req.Password, req.Role)
	if err != nil {
		respond(w, http.StatusConflict, dto.ErrorResponse{Error: "login already exists"})
		return
	}

	respond(w, 201, dto.OkResp{Ok: true})
}
