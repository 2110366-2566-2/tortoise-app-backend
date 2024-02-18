package services

import (
	"github.com/2110366-2566-2/tortoise-app-backend/internal/database"
	"github.com/2110366-2566-2/tortoise-app-backend/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/dgrijalva/jwt-go"
	"time"
)

func LoginHandler(c *gin.Context, h *database.Handler) {
	var loginRequest struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	user, err := database.ValidateLogin(c, h, loginRequest.Username, loginRequest.Password)

	if err != nil {
		c.JSON(401, gin.H{"error": err.Error()})
		return
	}

	var role string
	var email string
	if userRole, ok := user.(*models.User); ok {
		if userRole.Role == 1 {
			role = "seller"
			email = userRole.Email
		} else if userRole.Role == 2 {
			role = "buyer"
			email = userRole.Email
		} else {
			c.JSON(401, gin.H{"error": "user role not found"})
			return
		}
	} else if adminRole, ok := user.(*models.Admin); ok {
		role = "admin"
		email = adminRole.Email
	} else {
		c.JSON(401, gin.H{"error": "invalid user type"})
		return
	}

	// Create a new token object, specifying signing method and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"role":  role,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte("secret"))
	if err != nil {
		c.JSON(500, gin.H{"error": "failed to generate token"})
		return
	}

	c.JSON(200, gin.H{"token": tokenString})
	// c.JSON(200, gin.H{"user": &user})
}
