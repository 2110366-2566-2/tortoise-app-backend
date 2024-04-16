package utils

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"net/mail"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

// HashPassword hashes a password using bcrypt
func HashPassword(password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		panic(err)
	}
	return string(hash)
}

// ComparePassword compares a hashed password with a plaintext password
func ComparePassword(hashedPassword string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
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

func CreateTokenString(userID primitive.ObjectID, username string, role string) (string, error) {
	// Create a new token object, specifying signing method and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID":   userID,
		"username": username,
		"role":     role,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte("secret"))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ValidateBase64Image(base64Image string) (*[]string, error) {
	// Check if the string is empty
	if base64Image == "" {
		return nil, errors.New("empty image data")
	}

	// Check if the string starts with the correct prefix
	if !strings.HasPrefix(base64Image, "data:image/") {
		return nil, errors.New("invalid image format")
	}

	// Split the base64 image string
	splitString := strings.Split(base64Image, ";base64,")
	if len(splitString) != 2 {
		return nil, errors.New("invalid base64 image string")
	}

	// Decode the Base64-encoded image data
	_, err := base64.StdEncoding.DecodeString(splitString[1])
	if err != nil {
		return nil, fmt.Errorf("error decoding image data: %v", err)
	}

	return &splitString, nil

}

func GenerateHTMLTemplate(code, p1, p2, p3 string) string {
	html := `<style>
	body {
		font-family: Arial, sans-serif;
		color: #333;
		font-size: 14px;
	}
	.container {
		max-width: 600px;
		margin: 0 auto;
		padding: 20px;
		border: 1px solid #ccc;
		border-radius: 5px;
		background-color: #f9f9f9;
	}
	p {
		margin-bottom: 15px;
	}
	strong {
		color: #0066ff;
	}
	.OTP {
		color: #413a38;
		font-size: 20px;
		font-weight: bold;
	}
	.OTPText {
		font-size: 20px;
	}
	em {
		font-weight: bolder;
		font-style: normal;
	}
	.warntxt {
		color: #ff0000;
		font-size: 15px;
		font-style: italic;
	}
	.hello {
		font-size: 17px;
		font-weight: bold;
	}
	.special {
		color: #0b8114;
	}
</style>
<div class="container">
	<p class="hello">Hello,</p>
	<p>` + p1 + `</p>
	<p>` + p2 + `</p>
	<p class="OTPtext">
		<strong>` + p3 + `: </strong>
		<span class="OTP"> ` + code + `</span>
	</p>
	<p class="warntxt">
		If you didn't request this, you can safely ignore this email.
	</p>
	<p>Best regards,<br />PetPal User Support Team</p>
</div>
`
	return html
}
