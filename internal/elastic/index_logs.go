package elastic

import (
	"bytes"
	"encoding/json"

	"go.uber.org/zap"
)

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
