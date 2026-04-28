package handlers

import (
	"RacingGuru/internal/auth"
	"RacingGuru/internal/models"
	"RacingGuru/internal/shared"
	"context"
	"errors"
	"net/http"
	"runtime/debug"
	"strings"
	"time"

	"github.com/go-chi/chi/v5/middleware"
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
			h.Logger.Warn("authorization failed",
				"method", r.Method,
				"path", r.URL.Path,
				"request_id", middleware.GetReqID(r.Context()),
				"reason", "invalid authorization header",
			)
			h.sendError(w, "unauthorized", "UNAUTHORIZED", http.StatusUnauthorized)
			return
		}

		tokenString := parts[1]

		user, err := auth.ParseToken(tokenString)
		if err != nil {
			h.Logger.Warn("authorization failed",
				"method", r.Method,
				"path", r.URL.Path,
				"request_id", middleware.GetReqID(r.Context()),
				"reason", "invalid token",
				"error", err,
			)
			if errors.Is(err, shared.ErrorTokenExpired) {
				h.sendError(w, "token expired", "TOKEN_EXPIRED", http.StatusUnauthorized)
				return
			}
			h.sendError(w, "unauthorized", "UNAUTHORIZED", http.StatusUnauthorized)
			return
		}

		h.Logger.Debug("authorization succeeded",
			"method", r.Method,
			"path", r.URL.Path,
			"request_id", middleware.GetReqID(r.Context()),
			"user_id", user.Id,
		)

		ctx := context.WithValue(r.Context(), UserContextKey, user)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func getUserFromContext(ctx context.Context) (models.User, bool) {
	user, ok := ctx.Value(UserContextKey).(models.User)
	return user, ok
}

func (h *Handler) requestLoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestID := middleware.GetReqID(r.Context())
		startedAt := time.Now()

		h.Logger.Debug("request started",
			"method", r.Method,
			"path", r.URL.Path,
			"request_id", requestID,
			"remote_addr", r.RemoteAddr,
		)

		wrapped := middleware.NewWrapResponseWriter(w, r.ProtoMajor)
		next.ServeHTTP(wrapped, r)

		status := wrapped.Status()
		if status == 0 {
			status = http.StatusOK
		}

		args := []any{
			"method", r.Method,
			"path", r.URL.Path,
			"status", status,
			"bytes", wrapped.BytesWritten(),
			"duration", time.Since(startedAt),
			"request_id", requestID,
		}

		switch {
		case status >= http.StatusInternalServerError:
			h.Logger.Error("request completed", args...)
		case status >= http.StatusBadRequest:
			h.Logger.Warn("request completed", args...)
		case r.URL.Path == "/":
			h.Logger.Debug("request completed", args...)
		default:
			h.Logger.Info("request completed", args...)
		}
	})
}

func (h *Handler) recovererMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if rec := recover(); rec != nil {
				h.Logger.Error("panic recovered",
					"method", r.Method,
					"path", r.URL.Path,
					"request_id", middleware.GetReqID(r.Context()),
					"panic", rec,
					"stack", string(debug.Stack()),
				)

				if !responseStarted(w) {
					writeErrorResponse(w, "internal server error", "INTERNAL_ERROR", http.StatusInternalServerError)
				}
			}
		}()

		next.ServeHTTP(w, r)
	})
}

func responseStarted(w http.ResponseWriter) bool {
	type statusWriter interface {
		Status() int
	}

	wrapped, ok := w.(statusWriter)
	if !ok {
		return false
	}

	return wrapped.Status() != 0
}
