package handlers

import (
	"encoding/json"
	"net/http"

	"go.uber.org/zap"
)

func (h *Handler) handleAddLogs(w http.ResponseWriter, r *http.Request) {
	var req struct {
		ParseType string                   `json:"parse_type"`
		Logs      []map[string]interface{} `json:"logs"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.log.Error("failed to decode request", zap.Error(err))
		h.respond(w, http.StatusBadRequest, "invalid request body")
		return
	}
	var normalizedLogs []map[string]interface{}
	for _, log := range req.Logs {

		if normalized, err := h.pr.Parse(log, req.ParseType); err != nil {
			h.log.Error("failed to parse request", zap.Error(err))
			h.respond(w, http.StatusBadRequest, "failed to parse log")
			return
		} else {
			normalizedLogs = append(normalizedLogs, normalized)
		}
	}

	if err := h.es.IndexLogs(normalizedLogs); err != nil {
		h.log.Error("failed to index log", zap.Error(err))
		h.respond(w, http.StatusInternalServerError, "failed to index log")
		return
	}

	h.respond(w, http.StatusCreated, okResp)
}
