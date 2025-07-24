package elastic

import (
	"bytes"
	"fmt"
	"io"
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

func TestCountLogs(t *testing.T) {
	var body []byte
	client := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		body, _ = io.ReadAll(r.Body)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{"count":3}`)
	})

	count, err := client.CountLogs(map[string]string{"level": "info"})
	if err != nil {
		t.Fatalf("CountLogs error: %v", err)
	}
	if count != 3 {
		t.Errorf("expected count 3, got %d", count)
	}
	if !bytes.Contains(body, []byte("level")) {
		t.Errorf("expected request body to contain filter")
	}
}

func TestGetLogs(t *testing.T) {
	var body []byte
	client := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		body, _ = io.ReadAll(r.Body)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{"hits":{"hits":[{"_source":{"msg":"a"}}]}}`)
	})

	logs, err := client.GetLogs(map[string]string{"foo": "bar"}, 1, 0)
	if err != nil {
		t.Fatalf("GetLogs error: %v", err)
	}
	if len(logs) != 1 || logs[0]["msg"] != "a" {
		t.Errorf("unexpected logs %v", logs)
	}
	if !bytes.Contains(body, []byte("foo")) {
		t.Errorf("expected request body to contain filter")
	}
}

func TestIndexLog(t *testing.T) {
	var body []byte
	client := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		body, _ = io.ReadAll(r.Body)
		w.WriteHeader(200)
	})
	entry := map[string]interface{}{"msg": "hello"}
	if err := client.IndexLog(entry); err != nil {
		t.Fatalf("IndexLog error: %v", err)
	}
	if entry["raw"] == nil {
		t.Errorf("raw field not added")
	}
	if !bytes.Contains(body, []byte("hello")) {
		t.Errorf("expected body to contain log data")
	}
}

func TestIndexLogs(t *testing.T) {
	var body []byte
	client := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		body, _ = io.ReadAll(r.Body)
		w.WriteHeader(200)
	})
	logs := []map[string]interface{}{{"a": 1}, {"b": 2}}
	if err := client.IndexLogs(logs); err != nil {
		t.Fatalf("IndexLogs error: %v", err)
	}
	if logs[0]["raw"] == nil || logs[1]["raw"] == nil {
		t.Errorf("raw field not added")
	}
	if len(body) == 0 {
		t.Errorf("no body sent")
	}
}
