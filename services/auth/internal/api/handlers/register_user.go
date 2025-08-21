package api

import (
	"encoding/json"
	"logs-collector/pkg/dto"
	"logs-collector/pkg/dto/auth"
	"net/http"
	"strings"
)

func (h *Handlers) Register(w http.ResponseWriter, r *http.Request) {
	var req auth_dto.RegisterReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil || req.Login == "" || len(req.Password) < 3 {
		respond(w, 400, dto.ErrorResponse{Error: "invalid payload"})
		return
	}
	_, err := h.store.CreateUser(r.Context(), strings.ToLower(req.Login), req.Password, auth_dto.RoleUser)
	if err != nil {
		respond(w, 409, dto.ErrorResponse{Error: "login already exists"})
		return
	}
	respond(w, 201, dto.OkResp{Ok: true})
}
