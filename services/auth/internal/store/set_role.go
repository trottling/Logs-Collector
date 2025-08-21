package store

import (
	"context"
	"logs-collector/services/auth/internal/models"
)

func (s *Store) SetRole(ctx context.Context, id, role string) (bool, error) {
	tx := s.db.WithContext(ctx).Model(&models.User{}).
		Where("id = ?", id).
		Update("role", role)
	return tx.RowsAffected > 0, tx.Error
}
