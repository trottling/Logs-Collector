package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"log_stash_lite/internal/api/dto"

	"github.com/golang-jwt/jwt/v5"
)

// HandleAuthToken
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @Summary Generate JWT token
// @Description Generates a JWT token by user ID
// @Tags auth
// @Accept json
// @Produce json
// @Param data query dto.AuthTokenRequest false "Auth params"
// @Success 200 {object} dto.AuthTokenResponse
// @Failure 400 string dto.ErrorResponse
// @Failure 500 string dto.ErrorResponse
// @Router /auth/token [get]
func (h *Handler) HandleAuthToken(w http.ResponseWriter, r *http.Request) {
	var req dto.AuthTokenRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil || req.UserID == "" {
		req.UserID = r.URL.Query().Get("user_id")
	}

	if req.UserID == "" {
		h.respond(w, http.StatusBadRequest, dto.ErrorResponse{Error: "user_id is required"})
		return
	}

	if h.cfg.JWTSecret == "" {
		h.respond(w, http.StatusInternalServerError, dto.ErrorResponse{Error: "Invalid JWT secret on server"})
		return
	}

	claims := jwt.MapClaims{
		"sub": req.UserID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(h.cfg.JWTSecret))
	if err != nil {
		h.respond(w, http.StatusInternalServerError, dto.ErrorResponse{Error: "failed to generate token"})
		return
	}
	h.respond(w, http.StatusOK, dto.AuthTokenResponse{Token: tokenString})
}
