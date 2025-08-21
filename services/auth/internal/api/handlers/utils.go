package api

import (
	"encoding/json"
	"fmt"
	mw "logs-collector/pkg/middleware"
	"net/http"
)

func getRole(r *http.Request) string {
	if v := r.Context().Value(mw.RoleCtxKey()); v != nil {
		if s, ok := v.(string); ok {
			return s
		}
	}
	return ""
}

func parseInt(s string, def int) int {
	var x int
	if _, err := fmt.Sscanf(s, "%d", &x); err != nil {
		return def
	}
	return x
}

func respond(w http.ResponseWriter, code int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_ = json.NewEncoder(w).Encode(v)
}
