package elastic

import (
	"bytes"
	"encoding/json"

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
