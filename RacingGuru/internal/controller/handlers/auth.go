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

func newAuthUserResponse(user models.User) AuthUserResponse {
	return AuthUserResponse{
		Name:  user.Name,
		Email: user.Email,
		Role:  user.Role,
	}
}

// Me godoc
// @Summary Current user profile
// @Tags auth
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Success 200 {object} AuthUserResponse
// @Failure 401 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /me [get]
func (h *Handler) Me(w http.ResponseWriter, r *http.Request) {
	user, ok := h.requireUser(w, r)
	if !ok {
		return
	}

	uc := auth.NewUserMeUseCase(h.Repo, user)
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	profile, err := uc.Run(ctx)
	if err != nil {
		h.sendErrorExpanded(w, r, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(newAuthUserResponse(profile))
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
		h.sendLoggedError(w, r, "invalid request", "INVALID_REQUEST", http.StatusBadRequest, err)
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
		h.sendErrorExpanded(w, r, err)
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
		h.sendLoggedError(w, r, "invalid request body", "INVALID_REQUEST_BODY", http.StatusBadRequest, err)
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
		h.sendErrorExpanded(w, r, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(tokenS{token})
}

// Register godoc
// @Summary User registration
// @Description Registers a new user. If a valid one-time `secret` is provided, registers the first admin instead.
// @Tags auth
// @Accept json
// @Produce json
// @Param request body RegisterRequest true "Registration payload. Optional secret promotes the first user to admin"
// @Success 201 {object} AuthUserResponse "Registered user"
// @Failure 400 {object} ErrorResponse
// @Failure 403 {object} ErrorResponse
// @Failure 409 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /register [post]
func (h *Handler) Register(w http.ResponseWriter, r *http.Request) {
	request := RegisterRequest{}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		h.sendLoggedError(w, r, "invalid request", "INVALID_REQUEST", http.StatusBadRequest, err)
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
	if request.Secret != "" {
		if err := h.AdminSecrets.Validate(ctx, h.Repo, request.Secret); err != nil {
			h.sendErrorExpanded(w, r, err)
			return
		}
		user.Role = models.RoleAdmin
	}
	user, err := uc.Run(ctx, user)
	if err != nil {
		h.sendErrorExpanded(w, r, err)
		return
	}
	if request.Secret != "" && user.Role == models.RoleAdmin {
		h.AdminSecrets.Consume(request.Secret)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(newAuthUserResponse(user))
}

// GenerateAdminSecret godoc
// @Summary Generate one-time admin secret
// @Description Generates a one-time secret that can be passed in `/register` to create the first admin user.
// @Tags auth
// @Accept json
// @Produce json
// @Success 201 {object} GenerateAdminSecretResponse
// @Failure 409 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /setup/admin-secret [post]
func (h *Handler) GenerateAdminSecret(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	secret, err := h.AdminSecrets.Generate(ctx, h.Repo)
	if err != nil {
		h.sendErrorExpanded(w, r, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(GenerateAdminSecretResponse{Secret: secret})
}
