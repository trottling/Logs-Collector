package config

import (
	"os"
	"time"
)

type Config struct {
	Addr         string
	DatabaseURL  string
	JWTSecret    string
	AccessTTL    time.Duration
	RefreshTTL   time.Duration
	RootLogin    string
	RootPassword string
}

func getEnv(k, def string) string {
	if v := os.Getenv(k); v != "" {
		return v
	}
	return def
}

func mustParseDuration(s, def string) time.Duration {
	d, err := time.ParseDuration(getEnv(s, def))
	if err != nil {
		d, _ = time.ParseDuration(def)
	}
	return d
}

func Load() Config {
	return Config{
		Addr:         getEnv("AUTH_ADDR", ":8001"),
		DatabaseURL:  getEnv("DATABASE_URL", "postgres://logs:logs@localhost:5432/logs?sslmode=disable"),
		JWTSecret:    getEnv("JWT_SECRET", "change_me"),
		AccessTTL:    mustParseDuration("ACCESS_TTL", "15m"),
		RefreshTTL:   mustParseDuration("REFRESH_TTL", "168h"),
		RootLogin:    getEnv("ROOT_LOGIN", "root"),
		RootPassword: getEnv("ROOT_PASSWORD", "root"),
	}
}
