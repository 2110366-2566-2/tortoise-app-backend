package services

import (
	"github.com/2110366-2566-2/tortoise-app-backend/internal/database"
	"github.com/2110366-2566-2/tortoise-app-backend/internal/models"
	"github.com/2110366-2566-2/tortoise-app-backend/internal/storage"
	"github.com/2110366-2566-2/tortoise-app-backend/pkg/utils"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// RegisterHandler is a function to handle register request
func RegisterHandler(c *gin.Context, h *database.Handler, storage *storage.Handler) {
	var user models.User

	// Bind the request body to user model
	if err := c.BindJSON(&user); err != nil {
		c.JSON(400, gin.H{"success": false, "error": "invalid request body"})
		return
	}

	// Sanitize the user input
	utils.UserSaniatize(&user)

	var role string
	if user.Role == 1 {
		role = "seller"
	} else if user.Role == 2 {
		role = "buyer"
	} else {
		c.JSON(400, gin.H{"success": false, "error": "Invalid role"})
		return
	}

	// Hash the password
	user.Password = utils.HashPassword(user.Password)

	// Check if the user already exists
	if exist, err := database.CheckUserExist(c, h, user); err != nil {
		c.JSON(500, gin.H{"error": "Internal server error", "success": false})
		return
	} else if exist {
		c.JSON(400, gin.H{"error": "Username already exists", "success": false})
		return
	}

	// Check if the email already exists
	if exist, err := database.CheckEmailExist(c, h, user); err != nil {
		c.JSON(500, gin.H{"success": false, "error": "Internal server error"})
		return
	} else if exist {
		c.JSON(400, gin.H{"success": false, "error": "Email already exists"})
		return
	}

	user.ID = primitive.NewObjectID()

	// Convert image from base64 to url and add image to storage
	if user.Image != "" {
		imageURL, err := storage.AddImage(c, user.ID.Hex(), "users", user.Image)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error(), "success": false})
			return
		}
		user.Image = imageURL
	}

	// Create the user
	err := database.CreateUser(c, h, user)
	// Create the user
	if err != nil {
		c.JSON(500, gin.H{"error": "Internal server error", "success": false})
		return
	}

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := utils.CreateTokenString(user.ID, user.Username, role)
	if err != nil {
		c.JSON(500, gin.H{"error": "failed to generate token", "success": false})
		return
	}

	user.License = ""

	c.JSON(200, gin.H{"success": true, "message": "User created successfully", "user": &user, "token": tokenString})
}
