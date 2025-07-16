package elastic

import (
	_ "embed"

	"github.com/elastic/go-elasticsearch/v9"
	"go.uber.org/zap"
)

//go:embed templates/logs_template.json
var LogsTemplate []byte

//go:embed templates/search_logs.json
var SearchLogsTemplate []byte

//go:embed templates/count_logs.json
var CountLogsTemplate []byte

//go:embed templates/index_log.json
var IndexLogTemplate []byte

//go:embed templates/index_logs.json
var IndexLogsTemplate []byte

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
