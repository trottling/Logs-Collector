package handlers

import (
	"net/http"

	"go.uber.org/zap"
)

func (h *Handler) handleGetLogs(w http.ResponseWriter, r *http.Request) {
	filters := make(map[string]string)
	for key, values := range r.URL.Query() {
		if len(values) > 0 {
			filters[key] = values[0]
		}
	}

	logs, err := h.es.GetLogs(filters)
	if err != nil {
		h.log.Error("failed to get logs", zap.Error(err))
		h.respond(w, http.StatusInternalServerError, map[string]string{"error": "failed to fetch logs"})
		return
	}

	h.respond(w, http.StatusOK, logs)
}
