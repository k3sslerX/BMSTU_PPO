package auth

import (
	"crypto/sha256"
	"encoding/hex"
)

func hashPassword(password string) string {
	sum := sha256.Sum256([]byte(password))
	return hex.EncodeToString(sum[:])
}
