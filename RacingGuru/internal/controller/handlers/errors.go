package handlers

import (
	"RacingGuru/internal/shared"
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5/middleware"
)

func (h *Handler) sendErrorExpanded(w http.ResponseWriter, r *http.Request, err error) {
	message, code, status := mapError(err)
	h.logRequestError(r, err, code, status)
	h.sendError(w, message, code, status)
}

func (h *Handler) sendLoggedError(w http.ResponseWriter, r *http.Request, message, code string, status int, err error) {
	h.logRequestError(r, err, code, status)
	h.sendError(w, message, code, status)
}

func mapError(err error) (message, code string, status int) {
	if errors.Is(err, shared.ErrorPermissionDenied) {
		return "invalid request", "INVALID_REQUEST", http.StatusForbidden
	}
	if errors.Is(err, shared.ErrorNotFound) {
		return "not found", "NOT_FOUND", http.StatusNotFound
	}
	if errors.Is(err, shared.ErrorUserAlreadyExists) {
		return "user already exists", "USER_ALREADY_EXISTS", http.StatusConflict
	}
	if errors.Is(err, shared.ErrorAdminAlreadyExists) {
		return "admin already exists", "ADMIN_ALREADY_EXISTS", http.StatusConflict
	}
	if errors.Is(err, shared.ErrorAdminSecretAlreadyIssued) {
		return "admin secret already issued", "ADMIN_SECRET_ALREADY_ISSUED", http.StatusConflict
	}
	if errors.Is(err, shared.ErrorInvalidAdminSecret) {
		return "invalid admin secret", "INVALID_ADMIN_SECRET", http.StatusForbidden
	}
	if errors.Is(err, shared.ErrorInvalidToken) {
		return "invalid token", "INVALID_TOKEN", http.StatusForbidden
	}
	if errors.Is(err, shared.ErrorIncorrectPassword) {
		return "invalid password", "INVALID_PASSWORD", http.StatusForbidden
	}
	if errors.Is(err, shared.ErrorInvalidData) {
		return "invalid data", "INVALID_DATA", http.StatusForbidden
	}
	return "internal server error", "INTERNAL_ERROR", http.StatusInternalServerError
}

func (h *Handler) logRequestError(r *http.Request, err error, code string, status int) {
	args := []any{
		"error", err,
		"code", code,
		"status", status,
	}

	if r != nil {
		args = append(args,
			"method", r.Method,
			"path", r.URL.Path,
			"request_id", middleware.GetReqID(r.Context()),
		)
	}

	if status >= http.StatusInternalServerError {
		h.Logger.Error("request failed", args...)
		return
	}

	h.Logger.Warn("request failed", args...)
}
