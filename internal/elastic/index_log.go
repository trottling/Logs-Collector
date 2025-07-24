package elastic

import (
	"bytes"
	"encoding/json"
	"fmt"

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
	body, err := json.Marshal(entry)
	if err != nil {
		return err
	}

	res, err := c.ES.Index("logs", bytes.NewReader(body))
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.IsError() {
		c.Log.Error("failed to index log", zap.String("status", res.Status()))
		return fmt.Errorf("index error: %s", res.Status())
	}

	c.Log.Info("log indexed")
	return nil
}
