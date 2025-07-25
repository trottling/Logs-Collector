package handlers

import (
	"net/http"

	"log_stash_lite/internal/api/dto"
	"log_stash_lite/internal/api/validation"

	"go.uber.org/zap"
)

// handleGetLogs returns logs with filters and limit
func (h *Handler) handleGetLogs(w http.ResponseWriter, r *http.Request) {
	var req dto.GetLogsRequest

	var q struct {
		Limit  int `schema:"limit"`
		Offset int `schema:"offset"`
	}

	if err := queryDecoder.Decode(&q, r.URL.Query()); err != nil {
		h.log.Error("failed to decode query", zap.Error(err))
		h.respond(w, http.StatusBadRequest, "invalid query")
		return
	}

	filters := make(map[string]string)
	for key, values := range r.URL.Query() {
		if len(values) > 0 && key != "limit" && key != "offset" {
			filters[key] = values[0]
		}
	}

	req.Filters = filters
	req.Limit = q.Limit
	req.Offset = q.Offset
	// Validate request
	if err := validation.Validate.Struct(&req); err != nil {
		h.log.Error("validation error", zap.Error(err))
		h.respond(w, http.StatusBadRequest, "validation error")
		return
	}

	// Get logs from elastic
	logs, err := h.es.GetLogs(req.Filters, req.Limit, req.Offset)
	if err != nil {
		h.log.Error("failed to get logs", zap.Error(err))
		h.respond(w, http.StatusInternalServerError, map[string]string{"error": "failed to fetch logs"})
		return
	}

	h.respond(w, http.StatusOK, dto.GetLogsResponse{Logs: logs, Count: len(logs)})
}
