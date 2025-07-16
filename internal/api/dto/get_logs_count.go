package dto

type GetLogsCountRequest struct {
	Filters map[string]string `json:"filters" validate:"required"`
}

type GetLogsCountResponse struct {
	Count int `json:"count"`
}
