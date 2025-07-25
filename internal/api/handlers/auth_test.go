package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAddLog_Unauthorized(t *testing.T) {
	_ = httptest.NewRequest(http.MethodPost, "/add_log", nil)
	_ = httptest.NewRecorder()
	// router.ServeHTTP(w, req)
	// assert 401
}

func TestAddLog_ValidJWT(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "/add_log", nil)
	req.Header.Set("Authorization", "Bearer validtoken")
	_ = httptest.NewRecorder()
	// router.ServeHTTP(w, req)
	// assert 201
}
