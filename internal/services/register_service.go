package services

import (
	"github.com/2110366-2566-2/tortoise-app-backend/internal/database"
	"github.com/2110366-2566-2/tortoise-app-backend/internal/models"
	"github.com/2110366-2566-2/tortoise-app-backend/pkg/utils"
	"github.com/gin-gonic/gin"
)

// RegisterHandler is a function to handle register request
func RegisterHandler(c *gin.Context, h *database.Handler) {
	var user models.User

	// Bind the request body to user model
	if err := c.BindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// Hash the password
	user.Password = utils.HashPassword(user.Password)

	// Check if the user already exists
	if exist, err := database.CheckUserExist(c, h, user); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	} else if exist {
		c.JSON(400, gin.H{"error": "Username already exists"})
		return
	}

	// Check if the email already exists
	if exist, err := database.CheckEmailExist(c, h, user); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	} else if exist {
		c.JSON(400, gin.H{"error": "Email already exists"})
		return
	}

	res_id, err := database.CreateUser(c, h, user)
	// Create the user
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	user.ID = res_id

	c.JSON(200, gin.H{"message": "User created successfully", "user": &user})
}
