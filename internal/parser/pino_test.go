package parser

import (
	"reflect"
	"testing"
)

func TestParsePino(t *testing.T) {
	input := map[string]interface{}{"msg": "hi", "time": 456, "level": "warn"}
	want := map[string]interface{}{
		"message":   "hi",
		"timestamp": 456,
		"level":     "warn",
		"raw":       input,
	}

	got, err := ParsePino(input)
	if err != nil {
		t.Fatalf("ParsePino returned error: %v", err)
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("ParsePino = %#v, want %#v", got, want)
	}
}
