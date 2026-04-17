package handlers

import (
	"RacingGuru/internal/shared"
	"errors"
	"net/http"
)

func (h *Handler) sendErrorExpanded(w http.ResponseWriter, err error) {
	if errors.Is(err, shared.ErrorPermissionDenied) {
		h.sendError(w, "invalid request", "INVALID_REQUEST", http.StatusForbidden)
		return
	}
	if errors.Is(err, shared.ErrorNotFound) {
		h.sendError(w, "not found", "NOT_FOUND", http.StatusNotFound)
		return
	}
	if errors.Is(err, shared.ErrorUserAlreadyExists) {
		h.sendError(w, "user already exists", "USER_ALREADY_EXISTS", http.StatusConflict)
		return
	}
	if errors.Is(err, shared.ErrorAdminAlreadyExists) {
		h.sendError(w, "admin already exists", "ADMIN_ALREADY_EXISTS", http.StatusConflict)
		return
	}
	if errors.Is(err, shared.ErrorAdminSecretAlreadyIssued) {
		h.sendError(w, "admin secret already issued", "ADMIN_SECRET_ALREADY_ISSUED", http.StatusConflict)
		return
	}
	if errors.Is(err, shared.ErrorInvalidAdminSecret) {
		h.sendError(w, "invalid admin secret", "INVALID_ADMIN_SECRET", http.StatusForbidden)
		return
	}
	if errors.Is(err, shared.ErrorInvalidToken) {
		h.sendError(w, "invalid token", "INVALID_TOKEN", http.StatusForbidden)
		return
	}
	if errors.Is(err, shared.ErrorIncorrectPassword) {
		h.sendError(w, "invalid password", "INVALID_PASSWORD", http.StatusForbidden)
		return
	}
	if errors.Is(err, shared.ErrorInvalidData) {
		h.sendError(w, "invalid data", "INVALID_DATA", http.StatusForbidden)
		return
	}
	h.sendError(w, "internal server error", "INTERNAL_ERROR", http.StatusInternalServerError)
	return
}
