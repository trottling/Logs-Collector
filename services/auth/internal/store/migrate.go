package store

import (
	"context"
	"logs-collector/services/auth/internal/models"
)

// Migrate применяет схему и включает расширение pgcrypto
func (s *Store) Migrate(ctx context.Context) error {
	// Расширение для gen_random_uuid()
	if err := s.db.Exec(`CREATE EXTENSION IF NOT EXISTS pgcrypto`).Error; err != nil {
		return err
	}
	return s.db.WithContext(ctx).AutoMigrate(&models.User{})
}
