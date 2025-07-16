package dto

type AddLogsRequest struct {
	ParseType string                   `json:"parse_type" validate:"required,oneof=default zap logrus pino"`
	Logs      []map[string]interface{} `json:"logs" validate:"required,min=1,dive,required"`
}

type AddLogsResponse struct {
	Status string `json:"status"`
	Error  string `json:"error,omitempty"`
}
