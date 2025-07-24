package config

import (
	"os"
	"testing"
)

func TestLoadDefaults(t *testing.T) {
	os.Clearenv()
	cfg := Load()
	if cfg.ListenAddr != ":8080" {
		t.Errorf("ListenAddr got %s", cfg.ListenAddr)
	}
	if cfg.ElasticURL != "http://localhost:9200" {
		t.Errorf("ElasticURL got %s", cfg.ElasticURL)
	}
	if cfg.ElasticUsername != "elastic" {
		t.Errorf("ElasticUsername got %s", cfg.ElasticUsername)
	}
	if cfg.ElasticPassword != "change_me" {
		t.Errorf("ElasticPassword got %s", cfg.ElasticPassword)
	}
	if cfg.LogLevel != "info" {
		t.Errorf("LogLevel got %s", cfg.LogLevel)
	}
}

func TestLoadFromEnv(t *testing.T) {
	os.Setenv("LISTEN_ADDR", "0.0.0.0:9999")
	os.Setenv("ELASTIC_URL", "http://es:9200")
	os.Setenv("ELASTIC_USERNAME", "user")
	os.Setenv("ELASTIC_PASSWORD", "pass")
	os.Setenv("LOG_LEVEL", "debug")
	defer os.Clearenv()

	cfg := Load()
	if cfg.ListenAddr != "0.0.0.0:9999" ||
		cfg.ElasticURL != "http://es:9200" ||
		cfg.ElasticUsername != "user" ||
		cfg.ElasticPassword != "pass" ||
		cfg.LogLevel != "debug" {
		t.Errorf("unexpected config: %+v", cfg)
	}
}
