package handlers

import (
	"encoding/json"
	"net/http"

	"log_stash_lite/internal/api/dto"
	"log_stash_lite/internal/api/validation"

	"go.uber.org/zap"
)

// handleAddLogs adds multiple log entries
func (h *Handler) handleAddLogs(w http.ResponseWriter, r *http.Request) {
	var req dto.AddLogsRequest

	// Decode request body
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.log.Error("failed to decode request", zap.Error(err))
		h.respond(w, http.StatusBadRequest, dto.ErrorResponse{Error: "invalid request body"})
		return
	}

	// Validate request
	if err := validation.Validate.Struct(&req); err != nil {
		h.log.Error("validation error", zap.Error(err))
		h.respond(w, http.StatusBadRequest, dto.ErrorResponse{Error: "validation error"})
		return
	}

	// Parse each log
	var normalizedLogs []map[string]interface{}
	for _, log := range req.Logs {
		if normalized, err := h.pr.Parse(log, req.ParseType); err != nil {
			h.log.Error("failed to parse request", zap.Error(err))
			h.respond(w, http.StatusBadRequest, dto.ErrorResponse{Error: "failed to parse log"})
			return
		} else {
			normalizedLogs = append(normalizedLogs, normalized)
		}
	}

	// Index logs in elastic
	if err := h.es.IndexLogs(normalizedLogs); err != nil {
		h.log.Error("failed to index log", zap.Error(err))
		h.respond(w, http.StatusInternalServerError, dto.ErrorResponse{Error: "failed to index log"})
		return
	}

	h.respond(w, http.StatusCreated, dto.AddLogsResponse{Status: "ok"})
}
