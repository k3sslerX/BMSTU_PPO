package auth

import (
	"RacingGuru/internal/models"
	"RacingGuru/internal/shared"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

const tokenTTL = 72 * time.Hour

type tokenClaims struct {
	UserID string `json:"user_id"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

func ValidateJWTSecretKey() error {
	_, err := jwtSecretKey()
	return err
}

func GenerateToken(user models.User) (string, error) {
	secretKey, err := jwtSecretKey()
	if err != nil {
		return "", err
	}
	userID := user.Id.String()
	role := string(user.Role)

	now := time.Now()
	claims := tokenClaims{
		UserID: userID,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(now.Add(tokenTTL)),
			IssuedAt:  jwt.NewNumericDate(now),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secretKey))
}

func ParseToken(tokenString string) (models.User, error) {
	secretKey, err := jwtSecretKey()
	if err != nil {
		return models.User{}, err
	}

	claims := &tokenClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secretKey), nil
	}, jwt.WithExpirationRequired())

	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return models.User{}, shared.ErrorTokenExpired
		}
		return models.User{}, shared.ErrorInvalidToken
	}

	if !token.Valid || claims.UserID == "" || claims.Role == "" {
		return models.User{}, shared.ErrorInvalidToken
	}

	parsedUserID, err := uuid.Parse(claims.UserID)
	if err != nil {
		return models.User{}, shared.ErrorInvalidToken
	}

	return models.User{Id: parsedUserID, Role: models.Role(claims.Role)}, nil
}

func jwtSecretKey() (string, error) {
	secretKey := os.Getenv("JWT_SECRET_KEY")
	if secretKey == "" {
		return "", fmt.Errorf("JWT_SECRET_KEY is not set")
	}

	return secretKey, nil
}
