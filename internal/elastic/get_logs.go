package elastic

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func (c *Client) GetLogs(filters map[string]string, limit int) ([]map[string]interface{}, error) {
	var must []map[string]interface{}

	for field, value := range filters {
		must = append(must, map[string]interface{}{
			"term": map[string]interface{}{
				field: value,
			},
		})
	}

	query := map[string]interface{}{
		"query": map[string]interface{}{
			"bool": map[string]interface{}{
				"must": must,
			},
		},
	}

	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		return nil, fmt.Errorf("failed to encode query: %w", err)
	}

	res, err := c.ES.Search(
		c.ES.Search.WithIndex("logs"),
		c.ES.Search.WithBody(&buf),
		c.ES.Search.WithTrackTotalHits(true),
		c.ES.Search.WithPretty(),
		c.ES.Search.WithSize(limit),
	)
	if err != nil {
		return nil, fmt.Errorf("elasticsearch search failed: %w", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		return nil, fmt.Errorf("elasticsearch error: %s", res.Status())
	}

	var r map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		return nil, fmt.Errorf("failed to decode search response: %w", err)
	}

	var hits []map[string]interface{}
	for _, hit := range r["hits"].(map[string]interface{})["hits"].([]interface{}) {
		hits = append(hits, hit.(map[string]interface{})["_source"].(map[string]interface{}))
	}

	return hits, nil
}
