package dto

type GetLogsRequest struct {
	Filters map[string]string `json:"filters" validate:"required"`
	Limit   int               `json:"limit,omitempty"`
}

type GetLogsResponse struct {
	Logs  []map[string]interface{} `json:"logs"`
	Count int                      `json:"count"`
}
