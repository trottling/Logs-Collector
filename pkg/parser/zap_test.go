package parser

import (
	"reflect"
	"testing"
)

func TestParseZap(t *testing.T) {
	input := map[string]interface{}{"msg": "hello", "ts": 123, "level": "info"}
	want := map[string]interface{}{
		"message":   "hello",
		"timestamp": 123,
		"level":     "info",
		"raw":       input,
	}

	got, err := ParseZap(input)
	if err != nil {
		t.Fatalf("ParseZap returned error: %v", err)
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("ParseZap = %#v, want %#v", got, want)
	}
}
