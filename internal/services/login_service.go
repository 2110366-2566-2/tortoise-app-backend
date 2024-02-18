package services

import (
	"github.com/2110366-2566-2/tortoise-app-backend/internal/database"
	"github.com/gin-gonic/gin"
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

	c.JSON(200, gin.H{"user": &user})
}
