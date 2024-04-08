package utils

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"net/mail"
	"strconv"
	"strings"
	"time"

	"github.com/2110366-2566-2/tortoise-app-backend/internal/models"
	"github.com/golang-jwt/jwt/v5"
	"github.com/microcosm-cc/bluemonday"
	"go.mongodb.org/mongo-driver/bson"
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

var p = bluemonday.UGCPolicy()

func GetIntQueryParam(param string) (int, error) {
	// prevent xss attack
	paramValue := p.Sanitize(param)
	if paramValue == "" {
		return 0, nil // No value provided, return default
	}
	intValue, err := strconv.Atoi(paramValue)
	if err != nil {
		return 0, err // Invalid value
	}

	if intValue < 0 {
		return 0, errors.New("invalid value")
	}

	return intValue, nil
}

func ValidateStringQueryParam(param string) (string, error) {
	// prevent xss attack
	paramValue := p.Sanitize(param)
	if paramValue == "" {
		return "", nil // No value provided, return default
	}

	return paramValue, nil
}

func ValidateArrayQueryParam(param []string) ([]string, error) {
	// prevent xss attack
	var sanitizedParams []string
	sanitizedParams = append(sanitizedParams, param...)
	return sanitizedParams, nil
}

func SanitizeString(data string) string {
	return p.Sanitize(data)
}

func PetSanitize(pet *models.Pet) {
	pet.Name = p.Sanitize(pet.Name)
	pet.Description = p.Sanitize(pet.Description)
	pet.Category = p.Sanitize(pet.Category)
	pet.Species = p.Sanitize(pet.Species)
	pet.Behavior = p.Sanitize(pet.Behavior)
	pet.Media = p.Sanitize(pet.Media)
	pet.Sex = p.Sanitize(pet.Sex)
	for i := range pet.Medical_records {
		pet.Medical_records[i].Description = p.Sanitize(pet.Medical_records[i].Description)
		pet.Medical_records[i].Medical_date = p.Sanitize(pet.Medical_records[i].Medical_date)
		pet.Medical_records[i].Medical_id = p.Sanitize(pet.Medical_records[i].Medical_id)
	}
}

func BsonSanitize(data *bson.M) {
	for key, value := range *data {
		switch value := value.(type) {
		case string:
			(*data)[key] = p.Sanitize(value)
		case bson.M:
			BsonSanitize(&value)
		}
	}
}

func UserSaniatize(user *models.User) {
	user.Username = p.Sanitize(user.Username)
	user.Email = p.Sanitize(user.Email)
	user.FirstName = p.Sanitize(user.FirstName)
	user.LastName = p.Sanitize(user.LastName)
	// user.Password = p.Sanitize(user.Password)
	user.Gender = p.Sanitize(user.Gender)
	user.PhoneNumber = p.Sanitize(user.PhoneNumber)
	user.Image = p.Sanitize(user.Image)
	user.Address.Province = p.Sanitize(user.Address.Province)
	user.Address.District = p.Sanitize(user.Address.District)
	user.Address.SubDistrict = p.Sanitize(user.Address.SubDistrict)
	user.Address.PostalCode = p.Sanitize(user.Address.PostalCode)
	user.Address.Street = p.Sanitize(user.Address.Street)
	user.Address.Building = p.Sanitize(user.Address.Building)
	user.Address.HouseNumber = p.Sanitize(user.Address.HouseNumber)
}
