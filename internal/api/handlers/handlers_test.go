package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"log_stash_lite/internal/elastic"
	"log_stash_lite/internal/parser"

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

func newHandler(t *testing.T, es *elastic.Client) *Handler {
	return NewHandler(zap.NewNop(), es, parser.New(zap.NewNop()))
}

func TestHandleAddLog(t *testing.T) {
	var body []byte
	es := newElastic(t, func(w http.ResponseWriter, r *http.Request) {
		body, _ = io.ReadAll(r.Body)
		w.WriteHeader(200)
	})
	h := newHandler(t, es)

	reqBody := map[string]interface{}{
		"parse_type": "default",
		"log":        map[string]interface{}{"foo": "bar"},
	}
	b, _ := json.Marshal(reqBody)
	r := httptest.NewRequest(http.MethodPost, "/add_log", bytes.NewReader(b))
	w := httptest.NewRecorder()

	h.handleAddLog(w, r)

	if w.Code != http.StatusCreated {
		t.Fatalf("status %d", w.Code)
	}
	if !bytes.Contains(body, []byte("foo")) {
		t.Errorf("es not called")
	}
}

func TestHandleAddLogs(t *testing.T) {
	var body []byte
	es := newElastic(t, func(w http.ResponseWriter, r *http.Request) {
		body, _ = io.ReadAll(r.Body)
		w.WriteHeader(200)
	})
	h := newHandler(t, es)

	reqBody := map[string]interface{}{
		"parse_type": "default",
		"logs":       []map[string]interface{}{{"a": 1}},
	}
	b, _ := json.Marshal(reqBody)
	r := httptest.NewRequest(http.MethodPost, "/add_logs", bytes.NewReader(b))
	w := httptest.NewRecorder()

	h.handleAddLogs(w, r)

	if w.Code != http.StatusCreated {
		t.Fatalf("status %d", w.Code)
	}
	if !bytes.Contains(body, []byte("a")) {
		t.Errorf("es not called")
	}
}

func TestHandleGetLogs(t *testing.T) {
	es := newElastic(t, func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, `{"hits":{"hits":[{"_source":{"msg":"x"}}]}}`)
	})
	h := newHandler(t, es)

	r := httptest.NewRequest(http.MethodGet, "/get_logs?limit=1", nil)
	w := httptest.NewRecorder()

	h.handleGetLogs(w, r)

	if w.Code != http.StatusOK {
		t.Fatalf("status %d", w.Code)
	}
	if !bytes.Contains(w.Body.Bytes(), []byte("msg")) {
		t.Errorf("missing log in response")
	}
}

func TestHandleGetLogsCount(t *testing.T) {
	es := newElastic(t, func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, `{"count":4}`)
	})
	h := newHandler(t, es)

	r := httptest.NewRequest(http.MethodGet, "/logs_stats?level=info", nil)
	w := httptest.NewRecorder()

	h.handleLogStats(w, r)

	if w.Code != http.StatusOK {
		t.Fatalf("status %d", w.Code)
	}
}
