package handlers

import (
	"RacingGuru/internal/auth"
	"RacingGuru/internal/models"
	"context"
	"net/http"
	"strings"
)

type contextKey string

const (
	UserContextKey contextKey = "user"
)

func (h *Handler) authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			next.ServeHTTP(w, r)
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			h.sendError(w, "unauthorized", "UNAUTHORIZED", http.StatusUnauthorized)
			return
		}

		tokenString := parts[1]

		user, err := auth.ParseToken(tokenString)
		if err != nil {
			h.sendError(w, "unauthorized", "UNAUTHORIZED", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), UserContextKey, user)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func getUserFromContext(ctx context.Context) (models.User, bool) {
	user, ok := ctx.Value(UserContextKey).(models.User)
	return user, ok
}
