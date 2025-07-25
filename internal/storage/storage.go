package storage

import "context"

// Storage defines operations for log storage backends.
type Storage interface {
	IndexLog(ctx context.Context, entry map[string]interface{}) error
	IndexLogs(ctx context.Context, entries []map[string]interface{}) error
	GetLogs(ctx context.Context, filters map[string]string, limit, offset int) ([]map[string]interface{}, error)
	CountLogs(ctx context.Context, filters map[string]string) (int, error)
}
