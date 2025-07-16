package elastic

import (
	"bytes"
	"encoding/json"

	"go.uber.org/zap"
)

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
