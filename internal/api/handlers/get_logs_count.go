package handlers

import (
	"net/http"

	"log_stash_lite/internal/api/dto"
	"log_stash_lite/internal/api/validation"

	"go.uber.org/zap"
)

// handleGetLogsCount returns only count of logs by filters
func (h *Handler) handleGetLogsCount(w http.ResponseWriter, r *http.Request) {
	var req dto.GetLogsCountRequest

	filters := make(map[string]string)
	// Parse query params
	for key, values := range r.URL.Query() {
		if len(values) > 0 {
			filters[key] = values[0]
		}
	}

	req.Filters = filters
	// Validate request
	if err := validation.Validate.Struct(&req); err != nil {
		h.log.Error("validation error", zap.Error(err))
		h.respond(w, http.StatusBadRequest, "validation error")
		return
	}

	// Get count from elastic
	count, err := h.es.CountLogs(req.Filters)
	if err != nil {
		h.log.Error("failed to get logs", zap.Error(err))
		h.respond(w, http.StatusInternalServerError, map[string]string{"error": "failed to fetch logs"})
		return
	}

	h.respond(w, http.StatusOK, dto.GetLogsCountResponse{Count: count})
}
