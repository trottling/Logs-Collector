package dto

type (
	HealthResponse struct {
		Status        string `json:"status"`
		Error         string `json:"error"`
		ElasticStatus string `json:"elastic_status"`
		SystemStatus  struct {
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
		} `json:"system_status"`
	}
)
