package api

import (
	"logs-collector/pkg/dto"
	"logs-collector/pkg/dto/auth"
	"net/http"

	jwtlib "github.com/golang-jwt/jwt/v5"
)

func (h *Handlers) Me(w http.ResponseWriter, r *http.Request) {
	claims, ok := r.Context().Value(ctxClaims{}).(jwtlib.MapClaims)
	if !ok {
		respond(w, 401, dto.ErrorResponse{Error: "unauthorized"})
		return
	}
	id, _ := claims["sub"].(string)
	user, err := h.store.FindByID(r.Context(), id)
	if err != nil {
		respond(w, 404, dto.ErrorResponse{Error: "not found"})
		return
	}
	respond(w, 200, auth_dto.MeResp{ID: user.ID, Login: user.Login, Role: user.Role})
}
