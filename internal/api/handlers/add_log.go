package handlers

import (
	"encoding/json"
	"net/http"

	"go.uber.org/zap"
)

func (h *Handler) handleAddLog(w http.ResponseWriter, r *http.Request) {
	var req struct {
		ParseType string                 `json:"parse_type"`
		Log       map[string]interface{} `json:"log"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.log.Error("failed to decode request", zap.Error(err))
		h.respond(w, http.StatusBadRequest, "invalid request body")
		return
	}

	normalized, err := h.pr.Parse(req.Log, req.ParseType)
	if err != nil {
		h.log.Error("failed to parse request", zap.Error(err))
		h.respond(w, http.StatusBadRequest, "failed to parse log")
		return
	}

	if err := h.es.IndexLog(normalized); err != nil {
		h.log.Error("failed to index log", zap.Error(err))
		h.respond(w, http.StatusInternalServerError, "failed to index log")
		return
	}

	h.respond(w, http.StatusCreated, map[string]string{"status": "ok"})
}
