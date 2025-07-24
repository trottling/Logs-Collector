package elastic

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"testing"
)

func TestCountLogs(t *testing.T) {
	var body []byte
	client := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		body, _ = io.ReadAll(r.Body)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{"count":3}`)
	})

	count, err := client.CountLogs(context.Background(), map[string]string{"level": "info"})
	if err != nil {
		t.Fatalf("CountLogs error: %v", err)
	}
	if count != 3 {
		t.Errorf("expected count 3, got %d", count)
	}
	if !bytes.Contains(body, []byte("level")) {
		t.Errorf("expected request body to contain filter")
	}
}
