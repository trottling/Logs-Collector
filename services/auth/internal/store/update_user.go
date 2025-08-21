package store

import (
	"context"
	"logs-collector/pkg/dto/auth"
	"logs-collector/services/auth/internal/models"
	"strings"
)

func (s *Store) UpdateUser(ctx context.Context, id string, req auth_dto.UserUpdateReq) (bool, error) {
	updates := map[string]any{}

	if req.Login != nil {
		updates["login"] = strings.ToLower(*req.Login)
	}

	if req.Password != nil {
		updates["password"] = *req.Password
	}

	if req.Role != nil {
		updates["role"] = *req.Role
	}

	if len(updates) == 0 {
		return false, ErrNoChanges
	}

	tx := s.db.WithContext(ctx).Model(&models.User{}).Where("id = ?", id).Updates(updates)
	return tx.RowsAffected > 0, tx.Error
}
