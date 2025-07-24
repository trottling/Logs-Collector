package elastic

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"testing"
)

func TestIndexLog(t *testing.T) {
	var body []byte
	client := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		body, _ = io.ReadAll(r.Body)
		w.WriteHeader(200)
	})
	entry := map[string]interface{}{"msg": "hello"}
	if err := client.IndexLog(context.Background(), entry); err != nil {
		t.Fatalf("IndexLog error: %v", err)
	}
	if entry["raw"] == nil {
		t.Errorf("raw field not added")
	}
	if !bytes.Contains(body, []byte("hello")) {
		t.Errorf("expected body to contain log data")
	}
}
