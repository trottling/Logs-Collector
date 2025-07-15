package model

type LogEntry struct {
	Timestamp string                 `json:"timestamp"`
	Level     string                 `json:"level"`
	Source    string                 `json:"source"`
	Message   string                 `json:"message"`
	Meta      map[string]interface{} `json:"meta,omitempty"`
}
