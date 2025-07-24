package parser

import (
	"reflect"
	"testing"

	"go.uber.org/zap"
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
