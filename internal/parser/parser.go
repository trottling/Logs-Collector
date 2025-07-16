package parser

import (
	"errors"
	"strings"

	"go.uber.org/zap"
)

type LogParser struct {
	log *zap.Logger
}

// New creates a new LogParser instance
func New(log *zap.Logger) *LogParser {
	return &LogParser{log: log}
}

// Parse parses a log entry according to logType
func (parser *LogParser) Parse(log map[string]interface{}, logType string) (map[string]interface{}, error) {
	var parseFunc func(map[string]interface{}) (map[string]interface{}, error)

	switch strings.ToLower(logType) {
	case "default":
		return log, nil
	case "zap":
		parseFunc = ParseZap
	case "logrus":
		parseFunc = ParseLogrus
	case "pino":
		parseFunc = ParsePino
	default:
		return nil, errors.New("unknown parser type: " + logType)
	}

	// Parse log
	parsed, err := parseFunc(log)
	if err != nil {
		return nil, err
	}

	// auto-flatten raw to root
	for k, v := range log {
		if _, exists := parsed[k]; !exists {
			parsed[k] = v
		}
	}

	return parsed, nil
}
