package utils

import (
	"golang.org/x/crypto/bcrypt"
	"crypto/rand"
	"encoding/hex"
	"net/mail"
)

// HashPassword hashes a password using bcrypt
func HashPassword(password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if (err != nil) {
		panic(err)
	}
	return string(hash)
}

// ValidateEmail checks if a string is a valid email address
func ValidateEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

// GenerateRandomString generates a random string of the given length
func GenerateRandomString(length int) (string, error) {
	bytes := make([]byte, length)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}
