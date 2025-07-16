package dto

type (
	AddLogRequest struct {
		ParseType string                 `json:"parse_type" validate:"required,oneof=default zap logrus pino"`
		Log       map[string]interface{} `json:"log" validate:"required"`
	}
	AddLogResponse struct {
		Status string `json:"status"`
		Error  string `json:"error,omitempty"`
	}
)
