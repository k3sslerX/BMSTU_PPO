package auth

import (
	"RacingGuru/internal/models"
	"RacingGuru/internal/shared"
	"errors"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func TestParseTokenRejectsExpiredToken(t *testing.T) {
	t.Setenv("JWT_SECRET_KEY", "test-jwt-secret")

	token := signedTestToken(t, tokenClaims{
		UserID: uuid.NewString(),
		Role:   string(models.RoleUser),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(-time.Minute)),
		},
	})

	_, err := ParseToken(token)
	if !errors.Is(err, shared.ErrorInvalidToken) {
		t.Fatalf("ParseToken() error = %v, want %v", err, shared.ErrorInvalidToken)
	}
}

func TestParseTokenRequiresExpiration(t *testing.T) {
	t.Setenv("JWT_SECRET_KEY", "test-jwt-secret")

	token := signedTestToken(t, tokenClaims{
		UserID: uuid.NewString(),
		Role:   string(models.RoleUser),
	})

	_, err := ParseToken(token)
	if !errors.Is(err, shared.ErrorInvalidToken) {
		t.Fatalf("ParseToken() error = %v, want %v", err, shared.ErrorInvalidToken)
	}
}

func signedTestToken(t *testing.T, claims tokenClaims) string {
	t.Helper()

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte("test-jwt-secret"))
	if err != nil {
		t.Fatalf("SignedString() error = %v", err)
	}

	return token
}
