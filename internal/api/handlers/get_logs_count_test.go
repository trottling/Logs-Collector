package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandleGetLogsCount(t *testing.T) {
	called := false
	var body []byte
	es := newElastic(t, func(w http.ResponseWriter, r *http.Request) {
		called = true
		body, _ = io.ReadAll(r.Body)
		fmt.Fprint(w, `{"count":3}`)
	})
	h := newHandler(t, es)

	r := httptest.NewRequest(http.MethodGet, "/get_logs_count?level=info", nil)
	w := httptest.NewRecorder()

	h.handleGetLogsCount(w, r)

	if w.Code != http.StatusOK {
		t.Fatalf("status %d", w.Code)
	}
	if !called {
		t.Errorf("es not called")
	}
	var resp map[string]int
	if err := json.NewDecoder(w.Body).Decode(&resp); err != nil {
		t.Fatalf("decode response: %v", err)
	}
	if resp["count"] != 3 {
		t.Errorf("count %d", resp["count"])
	}
	if !bytes.Contains(body, []byte("level")) {
		t.Errorf("elastic not called with filters")
	}
}

func TestHandleGetLogsCount_ElasticError(t *testing.T) {
	es := newElastic(t, func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})
	h := newHandler(t, es)

	r := httptest.NewRequest(http.MethodGet, "/get_logs_count?level=debug", nil)
	w := httptest.NewRecorder()

	h.handleGetLogsCount(w, r)

	if w.Code != http.StatusInternalServerError {
		t.Fatalf("status %d", w.Code)
	}
}
