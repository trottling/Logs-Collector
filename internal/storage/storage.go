package storage

// Storage defines operations for log storage backends.
type Storage interface {
	IndexLog(entry map[string]interface{}) error
	IndexLogs(entries []map[string]interface{}) error
	GetLogs(filters map[string]string, limit, offset int) ([]map[string]interface{}, error)
	CountLogs(filters map[string]string) (int, error)
}
