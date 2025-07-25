package elastic

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"

	"github.com/elastic/go-elasticsearch/v9/esapi"
)

// CountLogs returns count of logs by filters
func (c *Client) CountLogs(ctx context.Context, filters map[string]string) (int, error) {
	var must []map[string]interface{}

	// Build must filter
	for field, value := range filters {
		must = append(must, map[string]interface{}{
			"term": map[string]interface{}{
				field: value,
			},
		})
	}

	mustJSON, err := json.Marshal(must)
	if err != nil {
		return 0, fmt.Errorf("failed to marshal filters: %w", err)
	}

	// Build query through template
	query := fmt.Sprintf(string(CountLogsTemplate), mustJSON)

	var buf bytes.Buffer
	buf.WriteString(query)

	// Count request with retry
	res, err := withRetry(ctx, func() (*esapi.Response, error) {
		return c.ES.Count(
			c.ES.Count.WithContext(ctx),
			c.ES.Count.WithIndex("logs"),
			c.ES.Count.WithBody(&buf),
		)
	})
	if err != nil {
		return 0, fmt.Errorf("elasticsearch count failed: %w", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		return 0, fmt.Errorf("elasticsearch error: %s", res.Status())
	}

	var r map[string]interface{}
	// Decode response
	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		return 0, fmt.Errorf("failed to decode count response: %w", err)
	}

	count, ok := r["count"].(float64)
	if !ok {
		return 0, fmt.Errorf("unexpected count type")
	}

	return int(count), nil
}
