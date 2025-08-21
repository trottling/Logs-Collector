package jwt

import (
	"errors"

	jwtlib "github.com/golang-jwt/jwt/v5"
)

type Manager struct {
	secret []byte
}

func New(secret string) *Manager {
	return &Manager{secret: []byte(secret)}
}

func (m *Manager) SignAccess(sub, role string) (string, error) {
	claims := jwtlib.MapClaims{
		"sub":  sub,
		"role": role,
	}
	token := jwtlib.NewWithClaims(jwtlib.SigningMethodHS256, claims)
	ss, err := token.SignedString(m.secret)
	return ss, err
}

func (m *Manager) Parse(token string) (jwtlib.MapClaims, error) {
	tok, err := jwtlib.Parse(token, func(t *jwtlib.Token) (any, error) {
		if t.Method.Alg() != jwtlib.SigningMethodHS256.Alg() {
			return nil, errors.New("invalid alg")
		}
		return m.secret, nil
	})
	if err != nil || !tok.Valid {
		return nil, errors.New("invalid token")
	}
	claims, ok := tok.Claims.(jwtlib.MapClaims)
	if !ok {
		return nil, errors.New("invalid claims")
	}
	return claims, nil
}

func (m *Manager) Secret() []byte { return m.secret }
