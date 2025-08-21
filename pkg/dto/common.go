package dto

type ErrorResponse struct {
	Error   string `json:"error"`
	Details string `json:"details"`
}

type OkResp struct {
	Ok bool `json:"ok"`
}

var OkRespRaw = []byte("ok")
