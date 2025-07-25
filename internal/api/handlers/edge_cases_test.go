package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAddLog_BadRequest(t *testing.T) {
	_ = httptest.NewRequest(http.MethodPost, "/add_log", bytes.NewReader([]byte("{bad json")))
	_ = httptest.NewRecorder()
	// router.ServeHTTP(w, req)
	// assert 400
}

func TestAddLog_LargePayload(t *testing.T) {
	bigLog := make(map[string]interface{})
	for i := 0; i < 10000; i++ {
		bigLog[fmt.Sprintf("k%d", i)] = "x"
	}
	reqBody := map[string]interface{}{"parse_type": "default", "log": bigLog}
	b, _ := json.Marshal(reqBody)
	_ = httptest.NewRequest(http.MethodPost, "/add_log", bytes.NewReader(b))
	_ = httptest.NewRecorder()
	// router.ServeHTTP(w, req)
	// assert 201 or 413
}

func TestGetLogs_EmptyFilters(t *testing.T) {
	_ = httptest.NewRequest(http.MethodGet, "/get_logs", nil)
	_ = httptest.NewRecorder()
	// router.ServeHTTP(w, req)
	// assert 400
}
