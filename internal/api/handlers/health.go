package handlers

import (
	"net/http"

	"logs-collector/internal/api/dto"
	"logs-collector/internal/health"
)

// HandleHealth
// @Summary Health check
// @Description Returns health status of the service and system
// @Tags health
// @Produce json
// @Success 200 {object} dto.HealthResponse
// @Failure 503 {object} dto.HealthResponse
// @Router /health [get]
func (h *Handler) HandleHealth(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	status, err := h.es.HealthCheck(ctx)
	if err != nil {
		h.respond(w, http.StatusServiceUnavailable, dto.HealthResponse{Status: "bad", Error: status.Error})
		return
	}

	// Get system stats
	sysHealth, err := health.GetSystemStats()
	if err != nil {
		h.respond(w, http.StatusServiceUnavailable, dto.HealthResponse{Status: "bad", Error: err.Error(), ElasticStatus: status.ElasticStatus})
		return
	}

	// Build response
	response := dto.HealthResponse{
		Status:        "ok",
		ElasticStatus: status.ElasticStatus,
		SystemStatus: struct {
			Cpu struct {
				UsagePercent float64 `json:"usage_percent"`
				Temperature  float64 `json:"temperature"`
			} `json:"cpu"`
			Ram struct {
				UsedMB  uint64 `json:"used_mb"`
				TotalMB uint64 `json:"total_mb"`
			} `json:"ram"`
			Rom struct {
				UsedMB  uint64 `json:"used_mb"`
				TotalMB uint64 `json:"total_mb"`
			}
		}{
			Cpu: struct {
				UsagePercent float64 `json:"usage_percent"`
				Temperature  float64 `json:"temperature"`
			}{
				UsagePercent: sysHealth.CPUUsagePercent,
				Temperature:  sysHealth.CPUTemp,
			},
			Ram: struct {
				UsedMB  uint64 `json:"used_mb"`
				TotalMB uint64 `json:"total_mb"`
			}{
				UsedMB:  sysHealth.RAMUsedMB,
				TotalMB: sysHealth.RAMTotalMB,
			},
			Rom: struct {
				UsedMB  uint64 `json:"used_mb"`
				TotalMB uint64 `json:"total_mb"`
			}{
				UsedMB:  sysHealth.DiskUsedMB,
				TotalMB: sysHealth.DiskTotalMB,
			},
		},
	}

	h.respond(w, http.StatusOK, response)
}
