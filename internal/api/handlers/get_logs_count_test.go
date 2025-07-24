package handlers

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandleGetLogsCountEndpoint(t *testing.T) {
	es := newElastic(t, func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, `{"count":5}`)
	})
	h := newHandler(t, es)

	r := httptest.NewRequest(http.MethodGet, "/count?level=info", nil)
	w := httptest.NewRecorder()

	h.handleGetLogsCount(w, r)

	if w.Code != http.StatusOK {
		t.Fatalf("status %d", w.Code)
	}
}
