package handlers

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandleGetLogsCount(t *testing.T) {
	es := newElastic(t, func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, `{"count":4}`)
	})
	h := newHandler(t, es)

	r := httptest.NewRequest(http.MethodGet, "/api/v1/logs_stats?level=info", nil)
	w := httptest.NewRecorder()

	h.handleLogStats(w, r)

	if w.Code != http.StatusOK {
		t.Fatalf("status %d", w.Code)
	}
}
