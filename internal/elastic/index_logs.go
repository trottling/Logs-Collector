package elastic

import (
	"bytes"
	"encoding/json"
	"fmt"

	"go.uber.org/zap"
)

// IndexLogs indexes multiple log entries in elasticsearch
func (c *Client) IndexLogs(entries []map[string]interface{}) error {
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

		res, err := c.ES.Index("logs", bytes.NewReader(data))
		if err != nil {
			return err
		}
		res.Body.Close()
		if res.IsError() {
			c.Log.Error("failed to index log", zap.String("status", res.Status()))
			return fmt.Errorf("index error: %s", res.Status())
		}
	}

	c.Log.Info("logs indexed")
	return nil
}
