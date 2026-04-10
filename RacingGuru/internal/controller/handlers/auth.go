package handlers

import (
	"RacingGuru/internal/core/auth"
	"RacingGuru/internal/models"
	"RacingGuru/internal/shared"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"time"
)

type tokenS struct {
	Token string `json:"token"`
}

// Login godoc
// @Summary User login
// @Tags auth
// @Accept json
// @Produce json
// @Param request body LoginRequest true "Login payload"
// @Success 200 {object} tokenS
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /login [post]
func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	user := models.User{}
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		h.sendError(w, "invalid request", "INVALID_REQUEST", http.StatusBadRequest)
		return
	}
	if user.Email == "" || user.Password == "" {
		h.sendError(w, "invalid request", "INVALID_REQUEST", http.StatusBadRequest)
		return
	}
	uc := auth.NewUserLoginUseCase(h.Repo)
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()
	token, err := uc.Run(ctx, user)
	if err != nil {
		if errors.Is(err, shared.ErrorIncorrectPassword) {
			h.sendError(w, "invalid request", "INVALID_REQUEST", http.StatusBadRequest)
			return
		}
		h.sendError(w, "internal server error", "INTERNAL__ERROR", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(tokenS{token})
}

// Register godoc
// @Summary User registration
// @Tags auth
// @Accept json
// @Produce json
// @Param request body RegisterRequest true "Registration payload"
// @Success 201 {object} models.User
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /register [post]
func (h *Handler) Register(w http.ResponseWriter, r *http.Request) {
	user := models.User{}
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		h.sendError(w, "invalid request", "INVALID_REQUEST", http.StatusBadRequest)
		return
	}
	if user.Email == "" || user.Password == "" || user.Role == "" {
		h.sendError(w, "invalid request body", "INVALID_REQUEST", http.StatusBadRequest)
		return
	}
	uc := auth.NewUserRegisterUseCase(h.Repo)
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()
	user, err := uc.Run(ctx, user)
	if err != nil {
		if errors.Is(err, shared.ErrorUserAlreadyExists) {
			h.sendError(w, "invalid request", "INVALID_REQUEST", http.StatusBadRequest)
			return
		}
		h.sendError(w, "internal server error", "INTERNAL_ERROR", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(user)
}
