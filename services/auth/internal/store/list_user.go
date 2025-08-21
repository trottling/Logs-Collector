package store

import (
	"context"
	"logs-collector/pkg/dto/auth"
	"logs-collector/services/auth/internal/models"
)

func (s *Store) ListUsers(ctx context.Context, q auth_dto.UserListReq) ([]auth_dto.UserResp, int64, error) {
	dbq := s.db.WithContext(ctx).Model(&models.User{})

	if q.LoginLike != "" {
		dbq = dbq.Where("login ILIKE ?", "%"+q.LoginLike+"%")
	}
	if q.Role != "" {
		dbq = dbq.Where("role = ?", q.Role)
	}

	// total
	var total int64
	if err := dbq.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// order whitelist
	orderBy := "created_at"
	switch q.OrderBy {
	case "login":
		orderBy = "login"
	case "role":
		orderBy = "role"
	case "created_at":
		orderBy = "created_at"
	}
	dir := "ASC"
	if q.Desc {
		dir = "DESC"
	}
	dbq = dbq.Order(orderBy + " " + dir)

	// page
	if q.Limit <= 0 || q.Limit > 100 {
		q.Limit = 20
	}
	if q.Offset < 0 {
		q.Offset = 0
	}
	var rows []models.User
	if err := dbq.Limit(q.Limit).Offset(q.Offset).
		Select("id, login, role").
		Find(&rows).Error; err != nil {
		return nil, 0, err
	}

	items := make([]auth_dto.UserResp, 0, len(rows))
	for _, r := range rows {
		items = append(items, auth_dto.UserResp{
			ID: r.ID, Login: r.Login, Role: r.Role,
		})
	}
	return items, total, nil
}
