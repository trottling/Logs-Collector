package handlers

import (
	"encoding/json"
	"fmt"
	handlers2 "logs-collector/internal/api/handlers"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandleGetLogs(t *testing.T) {

	called := false
	es := handlers2.newElastic(t, func(w http.ResponseWriter, r *http.Request) {
		called = true
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{"hits":{"hits":[{"_source":{"msg":"hello"}}]}}`)
	})

	h := handlers2.newHandler(t, es)
	r := httptest.NewRequest(http.MethodGet, "/get_logs?level=info&limit=1", nil)
	w := httptest.NewRecorder()

	h.HandleGetLogs(w, r)
	if w.Code != http.StatusOK {
		t.Fatalf("status %d", w.Code)
	}
	if !called {
		t.Errorf("es not called")
	}

	var resp struct {
		Logs  []map[string]interface{} `json:"logs"`
		Count int                      `json:"count"`
	}

	if err := json.NewDecoder(w.Body).Decode(&resp); err != nil {
		t.Fatalf("decode response: %v", err)
	}

	if resp.Count != 1 {
		t.Errorf("count %d", resp.Count)
	}

	if len(resp.Logs) == 0 || resp.Logs[0]["msg"] != "hello" {
		t.Errorf("unexpected logs %+v", resp.Logs)
	}
}

func TestHandleGetLogs_ElasticError(t *testing.T) {
	es := handlers2.newElastic(t, func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	h := handlers2.newHandler(t, es)
	r := httptest.NewRequest(http.MethodGet, "/get_logs?level=debug", nil)
	w := httptest.NewRecorder()

	h.HandleGetLogs(w, r)
	if w.Code != http.StatusInternalServerError {
		t.Fatalf("status %d", w.Code)
	}
}
