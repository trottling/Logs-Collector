package config

import "os"

type Config struct {
	ListenAddr string
	ElasticURL string
}

func Load() Config {
	return Config{
		ListenAddr: getEnv("LISTEN_ADDR", "8080"),
		ElasticURL: getEnv("ELASTIC_URL", "http://localhost:9200"),
	}
}

func getEnv(key, fallback string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return fallback
}
