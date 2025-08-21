package parser

import (
	"reflect"
	"testing"

	"go.uber.org/zap"
)

func TestLogParserParse(t *testing.T) {
	logger := zap.NewNop()
	pr := New(logger)

	input := map[string]interface{}{"msg": "z", "foo": "bar", "level": "info"}
	want := map[string]interface{}{
		"message": "z",
		"level":   "info",
		"raw":     input,
		"foo":     "bar",
		"msg":     "z",
	}

	got, err := pr.Parse(input, "zap")
	if err != nil {
		t.Fatalf("Parse returned error: %v", err)
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Parse = %#v, want %#v", got, want)
	}
}

func TestLogParserUnknownType(t *testing.T) {
	logger := zap.NewNop()
	pr := New(logger)

	_, err := pr.Parse(map[string]interface{}{}, "unknown")
	if err == nil {
		t.Fatalf("expected error for unknown parser type")
	}
}

func TestLogParserDefault(t *testing.T) {
	logger := zap.NewNop()
	pr := New(logger)

	input := map[string]interface{}{"foo": "bar"}

	got, err := pr.Parse(input, "default")
	if err != nil {
		t.Fatalf("Parse returned error: %v", err)
	}
	if !reflect.DeepEqual(got, input) {
		t.Errorf("Parse default = %#v, want %#v", got, input)
	}
}
