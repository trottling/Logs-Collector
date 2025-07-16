package elastic

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// GetLogs returns logs from elasticsearch by filters, limit and offset
func (c *Client) GetLogs(filters map[string]string, limit int, offset int) ([]map[string]interface{}, error) {
	var must []map[string]interface{}

	// Build must filters
	for field, value := range filters {
		must = append(must, map[string]interface{}{
			"term": map[string]interface{}{
				field: value,
			},
		})
	}

	mustJSON, err := json.Marshal(must)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal filters: %w", err)
	}

	// Create query with offset and limit
	query := fmt.Sprintf(`{"query":{"bool":{"must":%s}},"size":%d,"from":%d}`, mustJSON, limit, offset)

	var buf bytes.Buffer
	buf.WriteString(query)

	// Search request
	res, err := c.ES.Search(
		c.ES.Search.WithIndex("logs"),
		c.ES.Search.WithBody(&buf),
		c.ES.Search.WithTrackTotalHits(true),
		c.ES.Search.WithPretty(),
	)
	if err != nil {
		return nil, fmt.Errorf("elasticsearch search failed: %w", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		return nil, fmt.Errorf("elasticsearch error: %s", res.Status())
	}

	var r map[string]interface{}
	// Decode response
	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		return nil, fmt.Errorf("failed to decode search response: %w", err)
	}

	var hits []map[string]interface{}
	// Extract hits
	for _, hit := range r["hits"].(map[string]interface{})["hits"].([]interface{}) {
		hits = append(hits, hit.(map[string]interface{})["_source"].(map[string]interface{}))
	}

	return hits, nil
}
