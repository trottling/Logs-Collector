package dto

// Request for getting logs
type GetLogsRequest struct {
	Filters map[string]string `json:"filters" validate:"required"`
	Limit   int               `json:"limit,omitempty"`
	Offset  int               `json:"offset,omitempty"`
}

// Response with logs and count
type GetLogsResponse struct {
	Logs  []map[string]interface{} `json:"logs"`
	Count int                      `json:"count"`
}
