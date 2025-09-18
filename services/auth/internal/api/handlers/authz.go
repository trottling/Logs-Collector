package api

import (
	"context"
	"errors"
	"logs-collector/pkg/dto"
	mw "logs-collector/pkg/middleware"
	"net/http"
	"strings"

	jwtlib "github.com/golang-jwt/jwt/v5"
)

type ctxClaims struct{}

func (h *Handlers) Authz(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		hdr := r.Header.Get("Authorization")
		parts := strings.SplitN(hdr, " ", 2)
		if hdr == "" || len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			respond(w, 401, dto.ErrorResponse{Error: "missing/invalid auth header"})
			return
		}

		tok, err := jwtlib.Parse(parts[1], func(t *jwtlib.Token) (any, error) {
			if t.Method.Alg() != jwtlib.SigningMethodHS256.Alg() {
				return nil, errors.New("alg")
			}
			return h.jwt.Secret(), nil
		})

		if err != nil || !tok.Valid {
			respond(w, 401, dto.ErrorResponse{Error: "invalid token"})
			return
		}

		claims := tok.Claims.(jwtlib.MapClaims)
		role, _ := claims["role"].(string)
		ctx := context.WithValue(r.Context(), ctxClaims{}, claims)
		ctx = context.WithValue(ctx, mw.RoleCtxKey(), role)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
