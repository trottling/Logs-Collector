package handlers

import "net/http"

func (h *Handler) handleHealth(w http.ResponseWriter, _ *http.Request) {
	res, err := h.es.ES.Info()
	if err != nil || res.StatusCode != http.StatusOK {
		h.respond(w, http.StatusServiceUnavailable, map[string]string{"status": "elasticsearch unavailable"})
		return
	}
	defer res.Body.Close()

	h.respond(w, http.StatusOK, okResp)
}
