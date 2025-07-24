package handlers

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

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
