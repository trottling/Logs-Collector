package handlers

import (
	"fmt"
	"log_stash_lite/internal/api/dto"
	"log_stash_lite/internal/health"
	"net/http"
)

func (h *Handler) handleHealth(w http.ResponseWriter, _ *http.Request) {
	res, err := h.es.ES.Info()
	if err != nil {
		h.respond(w, http.StatusServiceUnavailable, dto.HealthResponse{Status: "bad", Error: fmt.Sprintf("Elastic health error: %s", err.Error())})
		return
	} else if res.StatusCode != http.StatusOK {
		h.respond(w, http.StatusServiceUnavailable, dto.HealthResponse{Status: "bad", Error: fmt.Sprintf("Bad elastic health status code: %d", res.StatusCode)})
		return
	}

	sysHealth, err := health.GetSystemStats()
	if err != nil {
		h.respond(w, http.StatusServiceUnavailable, dto.HealthResponse{Status: "bad", Error: err.Error(), ElasticStatus: "ok"})
		return
	}

	response := dto.HealthResponse{
		Status:        "ok",
		ElasticStatus: "ok",
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
