package api

import "net/http"

func (h *Handlers) Health(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(200)
	_, _ = w.Write([]byte("ok"))
}
