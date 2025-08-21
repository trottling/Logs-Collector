package elastic

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/elastic/go-elasticsearch/v9"
	"go.uber.org/zap"
)

func newTestClient(t *testing.T, handler http.HandlerFunc) *Client {
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
		t.Fatalf("failed to create es client: %v", err)
	}
	return &Client{ES: es, Log: zap.NewNop()}
}
