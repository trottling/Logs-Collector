package elastic

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"strings"

	"go.uber.org/zap"
)

// IndexLogs indexes multiple log entries in elasticsearch
func (c *Client) IndexLogs(entries []map[string]interface{}) error {
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

	c.Log.Debug("logs indexed")
	return nil
}
