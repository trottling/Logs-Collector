package elastic

import (
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
