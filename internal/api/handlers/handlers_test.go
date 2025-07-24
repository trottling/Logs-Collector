package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"log_stash_lite/internal/elastic"
	"log_stash_lite/internal/parser"
	"log_stash_lite/internal/storage"

	"github.com/elastic/go-elasticsearch/v9"
	"go.uber.org/zap"
)

func newElastic(t *testing.T, handler http.HandlerFunc) *elastic.Client {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-Elastic-Product", "Elasticsearch")
		handler(w, r)
	}))
	t.Cleanup(server.Close)
	es, err := elasticsearch.NewClient(elasticsearch.Config{
		Addresses: []string{server.URL},
		Transport: server.Client().Transport,
	})
	if err != nil {
		t.Fatalf("es client: %v", err)
	}
	return &elastic.Client{ES: es, Log: zap.NewNop()}
}

func newHandler(t *testing.T, es storage.Storage) *Handler {
	return NewHandler(zap.NewNop(), es, parser.New(zap.NewNop()))
}
