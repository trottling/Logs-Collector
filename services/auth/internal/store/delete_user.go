package store

import (
	"context"
	"logs-collector/services/auth/internal/models"
)

func (s *Store) DeleteUser(ctx context.Context, id string) (bool, error) {
	tx := s.db.WithContext(ctx).Delete(&models.User{}, "id = ?", id)
	return tx.RowsAffected > 0, tx.Error
}
