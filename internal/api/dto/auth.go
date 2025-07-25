package dto

type AuthTokenRequest struct {
	UserID string `json:"user_id" query:"user_id"`
}

type AuthTokenResponse struct {
	Token string `json:"token"`
}
