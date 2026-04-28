package handlers

import (
	"RacingGuru/internal/models"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func TestAuthMiddlewareReturnsTokenExpired(t *testing.T) {
	t.Setenv("JWT_SECRET_KEY", "test-jwt-secret")

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": uuid.NewString(),
		"role":    string(models.RoleUser),
		"exp":     time.Now().Add(-time.Minute).Unix(),
	}).SignedString([]byte("test-jwt-secret"))
	if err != nil {
		t.Fatalf("SignedString() error = %v", err)
	}

	handler := newTestHandler(t)
	request := httptest.NewRequest(http.MethodGet, "/sudoku/drivers", nil)
	request.Header.Set("Authorization", "Bearer "+token)
	response := httptest.NewRecorder()

	handler.authMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t.Fatal("next handler should not be called for expired token")
	})).ServeHTTP(response, request)

	if response.Code != http.StatusUnauthorized {
		t.Fatalf("authMiddleware() status = %d, want %d", response.Code, http.StatusUnauthorized)
	}
	if !strings.Contains(response.Body.String(), "TOKEN_EXPIRED") {
		t.Fatalf("authMiddleware() body = %q, want TOKEN_EXPIRED", response.Body.String())
	}
}
