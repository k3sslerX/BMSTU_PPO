package auth

import (
	"RacingGuru/internal/models"
	"RacingGuru/internal/shared"
	"fmt"
	"os"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(userID, role string) (string, error) {
	//secretKey := os.Getenv("JWT_SECRET_KEY")
	secretKey := "secretKey"

	claims := jwt.MapClaims{
		"user_id": userID,
		"role":    role,
		//"exp":     time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secretKey))
}

func ParseToken(tokenString string) (models.User, error) {
	secretKey := []byte(os.Getenv("JWT_SECRET_KEY"))
	//secretKey := []byte("secretKey")

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secretKey, nil
	})

	if err != nil {
		return models.User{}, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if userID, ok := claims["user_id"].(string); ok {
			if role, ok := claims["role"].(string); ok {
				return models.User{Id: userID, Role: models.Role(role)}, nil
			}
		}
	}

	return models.User{}, shared.ErrorInvalidToken
}
