package dto

// ErrorResponse is a common structure for API errors.

type ErrorResponse struct {
	Error string `json:"error"`
	Code  string `json:"code,omitempty"`
}
