package auth

import (
	"RacingGuru/internal/models"
	"RacingGuru/internal/shared"
	"fmt"
	"os"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

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

	claims := jwt.MapClaims{
		"user_id": userID,
		"role":    role,
		//"exp":     time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secretKey))
}

func ParseToken(tokenString string) (models.User, error) {
	secretKey, err := jwtSecretKey()
	if err != nil {
		return models.User{}, err
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secretKey), nil
	})

	if err != nil {
		return models.User{}, shared.ErrorInvalidToken
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if userID, ok := claims["user_id"].(string); ok {
			if role, ok := claims["role"].(string); ok {
				parsedUserID, err := uuid.Parse(userID)
				if err != nil {
					return models.User{}, shared.ErrorInvalidToken
				}
				return models.User{Id: parsedUserID, Role: models.Role(role)}, nil
			}
		}
	}

	return models.User{}, shared.ErrorInvalidToken
}

func jwtSecretKey() (string, error) {
	secretKey := os.Getenv("JWT_SECRET_KEY")
	if secretKey == "" {
		return "", fmt.Errorf("JWT_SECRET_KEY is not set")
	}

	return secretKey, nil
}
