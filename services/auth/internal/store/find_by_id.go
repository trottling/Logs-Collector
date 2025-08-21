package store

import (
	"context"
	"errors"
	"logs-collector/services/auth/internal/models"

	"gorm.io/gorm"
)

func (s *Store) FindByID(ctx context.Context, id string) (*models.User, error) {
	var u models.User
	err := s.db.WithContext(ctx).First(&u, "id = ?", id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, ErrNotFound
	}
	return &u, err
}
