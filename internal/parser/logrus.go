package parser

func ParseLogrus(log map[string]interface{}) (map[string]interface{}, error) {
	parsed := make(map[string]interface{})

	if msg, ok := log["message"]; ok {
		parsed["message"] = msg
	}
	if lvl, ok := log["level"]; ok {
		parsed["level"] = lvl
	}
	if time, ok := log["time"]; ok {
		parsed["timestamp"] = time
	}

	parsed["raw"] = log
	return parsed, nil
}
