package store

import (
	"context"
	"errors"
	"logs-collector/services/auth/internal/models"
	"strings"

	"gorm.io/gorm"
)

func (s *Store) FindByLogin(ctx context.Context, login string) (*models.User, error) {
	var u models.User
	err := s.db.WithContext(ctx).Where("login = ?", strings.ToLower(login)).First(&u).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, ErrNotFound
	}

	return &u, err
}
