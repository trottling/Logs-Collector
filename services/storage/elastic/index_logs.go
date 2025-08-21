package elastic

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/elastic/go-elasticsearch/v9/esapi"
	"go.uber.org/zap"
)

// IndexLogs indexes multiple log entries in elasticsearch
func (c *Client) IndexLogs(ctx context.Context, entries []map[string]interface{}) error {
	var entriesData []string
	for _, entry := range entries {
		// Marshal log entry to JSON
		rawData, err := json.Marshal(entry)
		if err != nil {
			return err
		}
		entry["raw"] = string(rawData)

		// Marshal entry with raw
		data, err := json.Marshal(entry)
		if err != nil {
			return err
		}
		entriesData = append(entriesData, string(data))
	}

	// Build query through template
	query := fmt.Sprintf(string(IndexLogsTemplate), strings.Join(entriesData, ",\n"))

	res, err := withRetry(ctx, func() (*esapi.Response, error) {
		return c.ES.API.Indices.Create(
			"logs",
			c.ES.API.Indices.Create.WithContext(ctx),
			c.ES.API.Indices.Create.WithBody(bytes.NewReader([]byte(query))),
		)
	})
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.IsError() {
		c.Log.Error("failed to index log", zap.String("status", res.Status()))
		return fmt.Errorf("elasticsearch error: %s", res.Status())
	}

	c.Log.Info("logs indexed")
	return nil
}
