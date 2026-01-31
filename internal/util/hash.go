package util

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
)

func HashPassword(password string) string {
	salt := make([]byte, 16)
	_, err := rand.Read(salt)
	if err != nil {
		return HashPasswordWithSalt(password, []byte{})
	}
	return HashPasswordWithSalt(password, salt)
}

func HashPasswordWithSalt(password string, salt []byte) string {
	h := sha256.New()
	h.Write(salt)
	h.Write([]byte(password))
	hashed := h.Sum(nil)

	result := make([]byte, len(salt)+len(hashed))
	copy(result, salt)
	copy(result[len(salt):], hashed)

	return hex.EncodeToString(result)
}
