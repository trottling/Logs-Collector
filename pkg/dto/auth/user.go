package auth_dto

// Роли
const (
	RoleRoot  = "root"
	RoleAdmin = "admin"
	RoleUser  = "user"
)

// UserCreateReq Создание пользователя
type UserCreateReq struct {
	Login    string `json:"login"    validate:"required,min=3,max=64,alphanum"`
	Password string `json:"password" validate:"required,min=3,max=128"` // сейчас plaintext; позже захэшируем
	Role     string `json:"role"     validate:"required,oneof=root admin user"`
}

type UserCreateResp struct {
	ID string `json:"id"`
}

// UserUpdateReq Частичное обновление пользователя
type UserUpdateReq struct {
	Login    *string `json:"login,omitempty"    validate:"omitempty,min=3,max=64,alphanum"`
	Password *string `json:"password,omitempty" validate:"omitempty,min=3,max=128"`
	Role     *string `json:"role,omitempty"     validate:"omitempty,oneof=root admin user"`
}

type UserUpdateResp struct {
	Updated bool `json:"updated"`
}

// UserDeleteResp Удаление пользователя
type UserDeleteResp struct {
	Deleted bool `json:"deleted"`
}

// UserResp Выдача одного пользователя
type UserResp struct {
	ID    string `json:"id"`
	Login string `json:"login"`
	Role  string `json:"role"`
}

// UserListReq Фильтры и листинг
type UserListReq struct {
	// Фильтры
	LoginLike string `json:"login_like,omitempty" validate:"omitempty,min=1,max=64"`
	Role      string `json:"role,omitempty"       validate:"omitempty,oneof=root admin user"`
	// Пагинация (offset/limit для простоты)
	Offset int `json:"offset" validate:"gte=0"`
	Limit  int `json:"limit"  validate:"gte=1,lte=100"`
	// Сортировка (минимализм)
	OrderBy string `json:"order_by" validate:"omitempty,oneof=login role created_at"`
	Desc    bool   `json:"desc"`
}

type UserListResp struct {
	Items []UserResp `json:"items"`
	Total int64      `json:"total"`
}
