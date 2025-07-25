package integration

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"logs-collector/internal/api/dto"
	"logs-collector/internal/api/handlers"
	"logs-collector/internal/config"
	"logs-collector/internal/elastic"
	"logs-collector/internal/parser"
	"logs-collector/internal/storage"
)

func TestAddLog_Integration(t *testing.T) {
	cfg := config.Load()
	log := parser.New(nil)
	es, err := elastic.NewClient(cfg, nil)
	if err != nil {
		t.Skip("Elastic not available: ", err)
	}
	var store storage.Storage = es
	h := handlers.NewHandler(nil, store, log, cfg)
	r := handlers.NewRouter(h, []byte(cfg.JWTSecret))

	reqBody := dto.AddLogRequest{
		ParseType: "default",
		Log:       map[string]interface{}{"foo": "bar"},
	}
	b, _ := json.Marshal(reqBody)
	req := httptest.NewRequest(http.MethodPost, "/add_log", bytes.NewReader(b))
	req.Header.Set("Authorization", "Bearer testtoken") // Put valid token here
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)
	if w.Code != http.StatusCreated {
		t.Fatalf("expected 201, got %d", w.Code)
	}
}
