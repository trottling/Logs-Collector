//go:build ignore

package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandleGetLogsCount(t *testing.T) {
	called := false
	es := newElastic(t, func(w http.ResponseWriter, r *http.Request) {
		called = true
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
