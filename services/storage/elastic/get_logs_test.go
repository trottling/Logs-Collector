package elastic

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"testing"
)

func TestGetLogs(t *testing.T) {
	var body []byte
	client := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		body, _ = io.ReadAll(r.Body)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{"hits":{"hits":[{"_source":{"msg":"a"}}]}}`)
	})

	logs, err := client.GetLogs(context.Background(), map[string]string{"foo": "bar"}, 1, 0)
	if err != nil {
		t.Fatalf("GetLogs error: %v", err)
	}
	if len(logs) != 1 || logs[0]["msg"] != "a" {
		t.Errorf("unexpected logs %v", logs)
	}
	if !bytes.Contains(body, []byte("foo")) {
		t.Errorf("expected request body to contain filter")
	}
}
