package handlers

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

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

func TestHandleAddLogs_ElasticError(t *testing.T) {
	es := newElastic(t, func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
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

	if w.Code != http.StatusInternalServerError {
		t.Fatalf("status %d", w.Code)
	}
}
