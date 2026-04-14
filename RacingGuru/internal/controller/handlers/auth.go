package handlers

import (
	"RacingGuru/internal/core/auth"
	"RacingGuru/internal/models"
	"context"
	"encoding/json"
	"net/http"
	"time"
)

type tokenS struct {
	Token string `json:"token"`
}

// ChangePassword godoc
// @Summary Change user password
// @Tags auth
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Param request body ChangePasswordRequest true "New password payload"
// @Success 204
// @Failure 400 {object} ErrorResponse
// @Failure 401 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /change-password [post]
func (h *Handler) ChangePassword(w http.ResponseWriter, r *http.Request) {
	user, ok := getUserFromContext(r.Context())
	if !ok {
		h.sendError(w, "unauthorized", "UNAUTHORIZED", http.StatusUnauthorized)
		return
	}

	request := ChangePasswordRequest{}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		h.sendError(w, "invalid request", "INVALID_REQUEST", http.StatusBadRequest)
		return
	}
	if request.Password == "" {
		h.sendError(w, "invalid request", "INVALID_REQUEST", http.StatusBadRequest)
		return
	}

	uc := auth.NewUserChangePasswordUseCase(h.Repo, user)
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	if err := uc.Run(ctx, request.Password); err != nil {
		h.sendErrorExpanded(w, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
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
		h.sendError(w, "invalid request body", "INVALID_REQUEST_BODY", http.StatusBadRequest)
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
		h.sendErrorExpanded(w, err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(tokenS{token})
}

// Register godoc
// @Summary User registration
// @Description Registers a new user with the default role `user`
// @Tags auth
// @Accept json
// @Produce json
// @Param request body RegisterRequest true "Registration payload without role"
// @Success 201 {object} models.User "Registered user with role user"
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /register [post]
func (h *Handler) Register(w http.ResponseWriter, r *http.Request) {
	request := RegisterRequest{}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		h.sendError(w, "invalid request", "INVALID_REQUEST", http.StatusBadRequest)
		return
	}
	if request.Email == "" || request.Password == "" {
		h.sendError(w, "invalid request body", "INVALID_REQUEST", http.StatusBadRequest)
		return
	}

	user := models.User{
		Name:     request.Name,
		Email:    request.Email,
		Password: request.Password,
		Role:     models.RoleUser,
	}

	uc := auth.NewUserRegisterUseCase(h.Repo)
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()
	user, err := uc.Run(ctx, user)
	if err != nil {
		h.sendErrorExpanded(w, err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(user)
}
