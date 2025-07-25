package handlers

import (
	"encoding/json"
	"net/http"

	"log_stash_lite/internal/api/dto"
	"log_stash_lite/internal/api/validation"

	"go.uber.org/zap"
)

// @Summary Add a log entry
// @Description Adds a single log entry to storage
// @Tags logs
// @Accept json
// @Produce json
// @Param data body dto.AddLogRequest true "Log entry"
// @Success 201 {object} dto.AddLogResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /add_log [post]
func (h *Handler) handleAddLog(w http.ResponseWriter, r *http.Request) {
	var req dto.AddLogRequest

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

	// Parse log
	normalized, err := h.pr.Parse(req.Log, req.ParseType)
	if err != nil {
		h.log.Error("failed to parse request", zap.Error(err))
		h.respond(w, http.StatusBadRequest, dto.ErrorResponse{Error: "failed to parse log"})
		return
	}

	// Index log in elastic
	if err := h.es.IndexLog(r.Context(), normalized); err != nil {
		h.log.Error("failed to index log", zap.Error(err))
		h.respond(w, http.StatusInternalServerError, dto.ErrorResponse{Error: "failed to index log"})
		return
	}

	h.respond(w, http.StatusCreated, dto.AddLogResponse{Status: "ok"})
}
