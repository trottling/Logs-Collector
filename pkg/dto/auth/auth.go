package auth_dto

// RegisterReq Регистрация
type RegisterReq struct {
	Login    string `json:"login"    validate:"required,min=3,max=64,alphanum"`
	Password string `json:"password" validate:"required,min=3,max=128"`
}

// RegisterResp Регистрация
type RegisterResp struct {
	ID string `json:"id"`
}

// LoginReq Логин
type LoginReq struct {
	Login    string `json:"login"    validate:"required"`
	Password string `json:"password" validate:"required"`
}

type TokenResp struct {
	AccessToken string `json:"access_token"`
}

type MeResp struct {
	ID    string `json:"id"`
	Login string `json:"login"`
	Role  string `json:"role"`
}
