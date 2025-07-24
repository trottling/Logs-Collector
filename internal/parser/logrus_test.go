package parser

import (
	"reflect"
	"testing"
)

func TestParseLogrus(t *testing.T) {
	input := map[string]interface{}{"message": "hello", "time": "2024-05-09T12:30:00Z", "level": "debug"}
	want := map[string]interface{}{
		"message":   "hello",
		"timestamp": "2024-05-09T12:30:00Z",
		"level":     "debug",
		"raw":       input,
	}

	got, err := ParseLogrus(input)
	if err != nil {
		t.Fatalf("ParseLogrus returned error: %v", err)
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("ParseLogrus = %#v, want %#v", got, want)
	}
}
