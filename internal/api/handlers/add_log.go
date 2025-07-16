package handlers

import (
	"encoding/json"
	"net/http"

	"log_stash_lite/internal/api/dto"
	"log_stash_lite/internal/api/validation"

	"go.uber.org/zap"
)

// handleAddLog adds a single log entry
func (h *Handler) handleAddLog(w http.ResponseWriter, r *http.Request) {
	var req dto.AddLogRequest

	// Decode request body
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.log.Error("failed to decode request", zap.Error(err))
		h.respond(w, http.StatusBadRequest, "invalid request body")
		return
	}

	// Validate request
	if err := validation.Validate.Struct(&req); err != nil {
		h.log.Error("validation error", zap.Error(err))
		h.respond(w, http.StatusBadRequest, "validation error")
		return
	}

	// Parse log
	normalized, err := h.pr.Parse(req.Log, req.ParseType)
	if err != nil {
		h.log.Error("failed to parse request", zap.Error(err))
		h.respond(w, http.StatusBadRequest, "failed to parse log")
		return
	}

	// Index log in elastic
	if err := h.es.IndexLog(normalized); err != nil {
		h.log.Error("failed to index log", zap.Error(err))
		h.respond(w, http.StatusInternalServerError, "failed to index log")
		return
	}

	h.respond(w, http.StatusCreated, dto.AddLogResponse{Status: "ok"})
}
