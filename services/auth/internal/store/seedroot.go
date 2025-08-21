package store

import (
	"logs-collector/pkg/dto/auth"
	"logs-collector/services/auth/internal/models"
	"strings"
)

// SeedRoot создаёт юзера root при отсутствии
func (s *Store) SeedRoot(login, password string) error {
	var count int64
	if err := s.db.Model(&models.User{}).Where("login = ?", strings.ToLower(login)).Count(&count).Error; err != nil {
		return err
	}

	if count > 0 {
		return nil
	}

	u := models.User{Login: strings.ToLower(login), Password: password, Role: auth_dto.RoleRoot}
	return s.db.Create(&u).Error
}
