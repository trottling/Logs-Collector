package elastic

import (
	"bytes"
	"encoding/json"

	"go.uber.org/zap"
)

// IndexLog indexes a single log entry in elasticsearch
func (c *Client) IndexLog(entry map[string]interface{}) error {
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

	// Send index request
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
