package services

import (
	"github.com/2110366-2566-2/tortoise-app-backend/pkg/utils"
	"github.com/gin-gonic/gin"
)

func LoginHandler(c *gin.Context) {
	var loginRequest struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	validCredentials := validateLogin(loginRequest.Username, loginRequest.Password)

	if !validCredentials {
		c.JSON(401, gin.H{"error": "Invalid credentials", "hint": "Arm is a pro golfer!"})
		return
	}

	c.JSON(200, gin.H{"message": "You got it!"})
}

func validateLogin(username string, password string) bool {
	hashedPassword := utils.HashPassword(password)
	if username != "arm" || hashedPassword != "766a08d07ad6a212c1e41f5efe975814d819bccd8c4a91ce81252820cd627e04" {
		return false
	}
	return true
}
