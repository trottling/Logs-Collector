package store

import (
	"context"
	"logs-collector/services/auth/internal/models"
	"strings"
)

func (s *Store) CreateUser(ctx context.Context, login, password, role string) (string, error) {
	u := models.User{
		Login: strings.ToLower(login), Password: password, Role: role,
	}

	if err := s.db.WithContext(ctx).Create(&u).Error; err != nil {
		return "", err
	}

	return u.ID, nil
}
