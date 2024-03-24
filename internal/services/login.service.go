package services

import (
	"fmt"
	"time"

	"github.com/2110366-2566-2/tortoise-app-backend/internal/database"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type loginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func LoginHandler(c *gin.Context, h *database.Handler) {
	var loginRequest loginRequest

	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	user, err := database.ValidateLogin(c, h, loginRequest.Username, loginRequest.Password)

	if err != nil {
		fmt.Println("Error: ", err)
		c.JSON(401, gin.H{"error": "invalid username or password"})
		return
	}

	var userID primitive.ObjectID
	var username string
	var role string

	if user.Role == 1 {
		role = "seller"
		userID = user.ID
		username = user.Username
	} else if user.Role == 2 {
		role = "buyer"
		userID = user.ID
		username = user.Username
	} else {
		fmt.Println("Error: invalid user type")
		c.JSON(401, gin.H{"error": "invalid user data"})
		return
	}

	tokenString, err := createTokenString(userID, username, role)
	if err != nil {
		fmt.Println("Error: ", err)
		c.JSON(500, gin.H{"error": "failed to generate token"})
		return
	}

	c.JSON(200, gin.H{"token": tokenString})
}

func LoginHandlerForAdmin(c *gin.Context, h *database.Handler) {
	var loginRequest loginRequest

	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	admin, err := database.AdminValidateLogin(c, h, loginRequest.Username, loginRequest.Password)

	if err != nil {
		fmt.Println("Error: ", err)
		c.JSON(401, gin.H{"error": "invalid username or password"})
		return
	}

	role := "admin"
	userID := admin.ID
	username := admin.Username

	tokenString, err := createTokenString(userID, username, role)
	if err != nil {
		fmt.Println("Error: ", err)
		c.JSON(500, gin.H{"error": "failed to generate token"})
		return
	}

	c.JSON(200, gin.H{"token": tokenString})
}

func createTokenString(userID primitive.ObjectID, username string, role string) (string, error) {
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
