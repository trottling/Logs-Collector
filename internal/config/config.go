package config

import "os"

// Config holds application configuration
type Config struct {
	ListenAddr string
	ElasticURL string
}

// Load loads config from environment variables
func Load() Config {
	return Config{
		ListenAddr: getEnv("LISTEN_ADDR", "8080"),
		ElasticURL: getEnv("ELASTIC_URL", "http://localhost:9200"),
	}
}

// getEnv returns environment variable or fallback
func getEnv(key, fallback string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return fallback
}
