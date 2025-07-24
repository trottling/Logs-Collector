package logger

import (
	"go.uber.org/zap"
)

// New creates a logger with the provided level string. Defaults to info on invalid values.
func New(level string) *zap.Logger {
	cfg := zap.NewProductionConfig()
	if err := cfg.Level.UnmarshalText([]byte(level)); err != nil {
		cfg.Level = zap.NewAtomicLevelAt(zap.InfoLevel)
	}
	logger, _ := cfg.Build()
	return logger
}
