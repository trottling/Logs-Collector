package middleware

import (
	"net/http"

	"logs-collector/pkg/dto/auth"
)

type ctxKey string

const roleKey ctxKey = "role"

func RequireRole(minRole string) func(http.Handler) http.Handler {
	rolePriority := map[string]int{
		auth_dto.RoleUser:  1,
		auth_dto.RoleAdmin: 2,
		auth_dto.RoleRoot:  3,
	}

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			role, _ := r.Context().Value(roleKey).(string)

			if role == "" {
				http.Error(w, "unauthorized", http.StatusUnauthorized)
				return
			}

			if rolePriority[role] < rolePriority[minRole] {
				http.Error(w, "forbidden", http.StatusForbidden)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}

// Экспортируй ключ, чтобы кладать/доставать роль в Authz
func RoleCtxKey() any { return roleKey }
