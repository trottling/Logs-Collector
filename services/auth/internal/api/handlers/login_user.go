package api

import (
	"crypto/subtle"
	"encoding/json"
	"logs-collector/pkg/dto"
	dtoAuth "logs-collector/pkg/dto/auth"
	"net/http"
	"strings"
)

func (h *Handlers) Login(w http.ResponseWriter, r *http.Request) {
	var req dtoAuth.LoginReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respond(w, http.StatusBadRequest, dto.ErrorResponse{Error: "invalid payload"})
		return
	}

	u, err := h.store.FindByLogin(r.Context(), strings.ToLower(req.Login))
	if err != nil {
		respond(w, http.StatusUnauthorized, dto.ErrorResponse{Error: "invalid credentials"})
		return
	}

	if subtle.ConstantTimeCompare([]byte(u.Password), []byte(req.Password)) != 1 {
		respond(w, http.StatusUnauthorized, dto.ErrorResponse{Error: "invalid credentials"})
		return
	}

	access, err := h.jwt.SignAccess(u.ID, u.Role)
	if err != nil {
		respond(w, http.StatusInternalServerError, dto.ErrorResponse{Error: "invalid credentials"})
		return
	}

	respond(w, 200, dtoAuth.TokenResp{AccessToken: access})
}
