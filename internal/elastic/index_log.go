package elastic

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"

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

	// Build query through template
	query := fmt.Sprintf(string(IndexLogTemplate), string(body))

	res, err := c.ES.API.Indices.Create("logs", c.ES.API.Indices.Create.WithBody(bytes.NewReader([]byte(query))))
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.IsError() {
		body, _ := io.ReadAll(res.Body)
		c.Log.Error("failed to index log", zap.String("status", res.Status()), zap.ByteString("response", body))
		return fmt.Errorf("elasticsearch error: %s", res.Status())
	}

	c.Log.Debug("log indexed")
	return nil
}
