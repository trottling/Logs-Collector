package parser

func ParseZap(log map[string]interface{}) (map[string]interface{}, error) {
	parsed := make(map[string]interface{})

	if msg, ok := log["msg"]; ok {
		parsed["message"] = msg
	}
	if ts, ok := log["ts"]; ok {
		parsed["timestamp"] = ts
	}
	if lvl, ok := log["level"]; ok {
		parsed["level"] = lvl
	}

	parsed["raw"] = log
	return parsed, nil
}
