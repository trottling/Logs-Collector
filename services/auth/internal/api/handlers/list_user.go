package api

import (
	"logs-collector/pkg/dto"
	"logs-collector/pkg/dto/auth"
	"net/http"
)

func (h *Handlers) ListUsers(w http.ResponseWriter, r *http.Request) {
	var req auth_dto.UserListReq
	req.LoginLike = r.URL.Query().Get("login_like")
	req.Role = r.URL.Query().Get("role")
	req.Offset = parseInt(r.URL.Query().Get("offset"), 0)
	req.Limit = parseInt(r.URL.Query().Get("limit"), 20)
	req.OrderBy = r.URL.Query().Get("order_by")
	if req.Limit <= 0 || req.Limit > 100 {
		req.Limit = 20
	}

	items, total, err := h.store.ListUsers(r.Context(), req)
	if err != nil {
		respond(w, 500, dto.ErrorResponse{Error: "db error"})
		return
	}
	respond(w, 200, auth_dto.UserListResp{Items: items, Total: total})
}
