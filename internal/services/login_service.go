package services

import (
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
	if username != "arm" || password != "proGolfer" {
		return false
	}
	return true
}
