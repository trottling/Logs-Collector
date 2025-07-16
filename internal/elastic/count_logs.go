package elastic

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// CountLogs returns count of logs by filters
func (c *Client) CountLogs(filters map[string]string) (int, error) {
	var must []map[string]interface{}

	// Build must filters
	for field, value := range filters {
		must = append(must, map[string]interface{}{
			"term": map[string]interface{}{
				field: value,
			},
		})
	}

	// Build query
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"bool": map[string]interface{}{
				"must": must,
			},
		},
	}

	var buf bytes.Buffer
	// Encode query to buffer
	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		return 0, fmt.Errorf("failed to encode query: %w", err)
	}

	// Count request
	res, err := c.ES.Count(
		c.ES.Count.WithIndex("logs"),
		c.ES.Count.WithBody(&buf),
	)
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
