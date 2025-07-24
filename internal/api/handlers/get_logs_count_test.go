package handlers

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandleGetLogsCountOK(t *testing.T) {
	es := newElastic(t, func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, `{"count":3}`)
	})
	h := newHandler(t, es)

	r := httptest.NewRequest(http.MethodGet, "/logs_count?level=info", nil)
	w := httptest.NewRecorder()

	h.handleGetLogsCount(w, r)

	if w.Code != http.StatusOK {
		t.Fatalf("status %d", w.Code)
	}
}
