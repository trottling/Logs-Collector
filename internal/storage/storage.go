package storage

import "context"

type HealthStatus struct {
	ElasticStatus string
	Error         string
}

type Storage interface {
	IndexLog(ctx context.Context, entry map[string]interface{}) error
	IndexLogs(ctx context.Context, entries []map[string]interface{}) error
	GetLogs(ctx context.Context, filters map[string]string, limit, offset int) ([]map[string]interface{}, error)
	CountLogs(ctx context.Context, filters map[string]string) (int, error)
	HealthCheck(ctx context.Context) (*HealthStatus, error)
}
