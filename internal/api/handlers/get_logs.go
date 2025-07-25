package handlers

import (
	"fmt"
	"net/http"

	"logs-collector/internal/api/dto"
	"logs-collector/internal/api/validation"

	"go.uber.org/zap"
)

// HandleGetLogs
// @Summary Get logs
// @Description Returns logs with filters and limit
// @Tags logs
// @Accept json
// @Produce json
// @Param level query string false "Log level"
// @Param limit query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {object} dto.GetLogsResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Security BearerAuth
// @Router /get_logs [get]
func (h *Handler) HandleGetLogs(w http.ResponseWriter, r *http.Request) {
	var req dto.GetLogsRequest

	filters := make(map[string]string)
	// Parse query params
	for key, values := range r.URL.Query() {
		if len(values) > 0 {
			filters[key] = values[0]
		}
	}

	// Parse limit param
	limit := 0
	if l, ok := filters["limit"]; ok {
		fmt.Sscanf(l, "%d", &limit)
		delete(filters, "limit")
	}

	// Parse offset param
	offset := 0
	if o, ok := filters["offset"]; ok {
		fmt.Sscanf(o, "%d", &offset)
		delete(filters, "offset")
	}

	req.Filters = filters
	req.Limit = limit
	req.Offset = offset

	// Validate request
	if err := validation.Validate.Struct(&req); err != nil {
		h.log.Error("validation error", zap.Error(err))
		h.respond(w, http.StatusBadRequest, dto.ErrorResponse{Error: "validation error"})
		return
	}

	// Get logs from elastic
	logs, err := h.es.GetLogs(r.Context(), req.Filters, req.Limit, req.Offset)
	if err != nil {
		h.log.Error("failed to get logs", zap.Error(err))
		h.respond(w, http.StatusInternalServerError, dto.ErrorResponse{Error: "failed to fetch logs"})
		return
	}

	h.respond(w, http.StatusOK, dto.GetLogsResponse{Logs: logs, Count: len(logs)})
}
