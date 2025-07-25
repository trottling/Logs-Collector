package handlers

import (
	"net/http"

	"logs-collector/internal/api/dto"
	"logs-collector/internal/api/validation"

	"go.uber.org/zap"
)

// HandleGetLogsCount returns only count of logs by filters
func (h *Handler) HandleGetLogsCountTest(w http.ResponseWriter, r *http.Request) {
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
		h.respond(w, http.StatusBadRequest, dto.ErrorResponse{Error: "validation error"})
		return
	}

	// Get count from elastic
	count, err := h.es.CountLogs(r.Context(), req.Filters)
	if err != nil {
		h.log.Error("failed to get logs", zap.Error(err))
		h.respond(w, http.StatusInternalServerError, dto.ErrorResponse{Error: "failed to fetch logs"})
		return
	}

	h.respond(w, http.StatusOK, dto.GetLogsCountResponse{Count: count})
}
