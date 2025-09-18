package config

import (
	"os"
)

type Config struct {
	Addr         string
	DatabaseURL  string
	LogLevel     string
	JWTSecret    string
	RootLogin    string
	RootPassword string
}

func getEnv(k, def string) string {
	if v := os.Getenv(k); v != "" {
		return v
	}
	return def
}

func Load() Config {
	return Config{
		Addr:         getEnv("AUTH_ADDR", ":8001"),
		DatabaseURL:  getEnv("DATABASE_URL", "postgres://logs:logs@localhost:5432/logs?sslmode=disable"),
		LogLevel:     getEnv("LOG_LEVEL", "1"),
		JWTSecret:    getEnv("JWT_SECRET", "change_me"),
		RootLogin:    getEnv("ROOT_LOGIN", "root"),
		RootPassword: getEnv("ROOT_PASSWORD", "root"),
	}
}
