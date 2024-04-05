package services

import (
	"fmt"

	"github.com/2110366-2566-2/tortoise-app-backend/internal/database"
	"github.com/2110366-2566-2/tortoise-app-backend/pkg/utils"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type loginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func LoginHandler(c *gin.Context, h *database.Handler) {
	var loginRequest loginRequest

	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		c.JSON(400, gin.H{"success": false, "error": "invalid request body"})
		return
	}

	user, err := database.ValidateLogin(c, h, loginRequest.Username, loginRequest.Password)

	if err != nil {
		fmt.Println("Error: ", err)
		c.JSON(401, gin.H{"success": false, "error": "invalid username or password"})
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
		c.JSON(401, gin.H{"success": false, "error": "invalid user data"})
		return
	}

	tokenString, err := utils.CreateTokenString(userID, username, role)
	if err != nil {
		fmt.Println("Error: ", err)
		c.JSON(500, gin.H{"success": false, "error": "failed to generate token"})
		return
	}

	c.JSON(200, gin.H{"success": true, "token": tokenString})
}
