package parser

func ParsePino(log map[string]interface{}) (map[string]interface{}, error) {
	parsed := make(map[string]interface{})

	if msg, ok := log["msg"]; ok {
		parsed["message"] = msg
	}
	if lvl, ok := log["level"]; ok {
		parsed["level"] = lvl
	}
	if ts, ok := log["time"]; ok {
		parsed["timestamp"] = ts
	}

	parsed["raw"] = log
	return parsed, nil
}
