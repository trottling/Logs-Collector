package elastic

import (
	"io"
	"net/http"
	"testing"
)

func TestIndexLogs(t *testing.T) {
	var body []byte
	client := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		body, _ = io.ReadAll(r.Body)
		w.WriteHeader(200)
	})
	logs := []map[string]interface{}{{"a": 1}, {"b": 2}}
	if err := client.IndexLogs(logs); err != nil {
		t.Fatalf("IndexLogs error: %v", err)
	}
	if logs[0]["raw"] == nil || logs[1]["raw"] == nil {
		t.Errorf("raw field not added")
	}
	if len(body) == 0 {
		t.Errorf("no body sent")
	}
}
