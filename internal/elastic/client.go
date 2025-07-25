package elastic

import (
	_ "embed"
	"log_stash_lite/internal/config"
	"log_stash_lite/internal/storage"

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

// Ensure Client implements the storage.Storage interface.
var _ storage.Storage = (*Client)(nil)

func NewClient(cfg config.Config, log *zap.Logger) (*Client, error) {
	es, err := elasticsearch.NewClient(elasticsearch.Config{
		Addresses: []string{cfg.ElasticURL},
		Username:  cfg.ElasticUsername,
		Password:  cfg.ElasticPassword,
	})
	if err != nil {
		return nil, err
	}
	return &Client{ES: es, Log: log}, nil
}
