package config

import "os"

// Config holds application configuration
type Config struct {
	ListenAddr      string
	ElasticURL      string
	ElasticUsername string
	ElasticPassword string
	LogLevel        string
}

// Load loads config from environment variables
func Load() Config {
	return Config{
		ListenAddr:      getEnv("LISTEN_ADDR", ":8080"),
		ElasticURL:      getEnv("ELASTIC_URL", "http://localhost:9200"),
		ElasticUsername: getEnv("ELASTIC_USERNAME", "elastic"),
		ElasticPassword: getEnv("ELASTIC_PASSWORD", "change_me"),
		LogLevel:        getEnv("LOG_LEVEL", "info"),
	}
}

// getEnv returns environment variable or fallback
func getEnv(key, fallback string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return fallback
}
