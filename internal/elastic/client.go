package elastic

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/elastic/go-elasticsearch/v9"
	"go.uber.org/zap"
)

type Client struct {
	ES  *elasticsearch.Client
	Log *zap.Logger
}

func NewClient(url string, log *zap.Logger) (*Client, error) {
	es, err := elasticsearch.NewClient(elasticsearch.Config{
		Addresses: []string{url},
	})
	if err != nil {
		return nil, err
	}
	return &Client{ES: es, Log: log}, nil
}

func (c *Client) IndexLog(entry map[string]interface{}) error {
	rawData, err := json.Marshal(entry)
	if err != nil {
		return err
	}
	entry["raw"] = string(rawData)

	data, err := json.Marshal(entry)
	if err != nil {
		return err
	}

	res, err := c.ES.Index("logs", bytes.NewReader(data))
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.IsError() {
		c.Log.Error("failed to index log", zap.String("status", res.Status()))
		return err
	}

	c.Log.Info("log indexed")
	return nil
}

func (c *Client) IndexLogs(entries []map[string]interface{}) error {
	for _, entry := range entries {
		rawData, err := json.Marshal(entry)
		if err != nil {
			return err
		}
		entry["raw"] = string(rawData)

		data, err := json.Marshal(entry)
		if err != nil {
			return err
		}

		res, err := c.ES.Index("logs", bytes.NewReader(data))
		if err != nil {
			return err
		}

		defer res.Body.Close()

		if res.IsError() {
			c.Log.Error("failed to index log", zap.String("status", res.Status()))
			return err
		}
	}

	c.Log.Info("logs indexed")
	return nil
}

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

func (c *Client) CountLogs(filters map[string]string) (int, error) {
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
		return 0, fmt.Errorf("failed to encode query: %w", err)
	}

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
	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		return 0, fmt.Errorf("failed to decode count response: %w", err)
	}

	count, ok := r["count"].(float64)
	if !ok {
		return 0, fmt.Errorf("unexpected count type")
	}

	return int(count), nil
}
