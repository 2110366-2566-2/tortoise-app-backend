package services

import (
	"fmt"
	"time"

	"github.com/2110366-2566-2/tortoise-app-backend/internal/database"
	"github.com/2110366-2566-2/tortoise-app-backend/internal/models"
	"github.com/2110366-2566-2/tortoise-app-backend/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson"
)

// RegisterHandler is a function to handle register request
func RegisterHandler(c *gin.Context, h *database.Handler) {
	var user models.User

	// Bind the request body to user model
	if err := c.BindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	var role string
	if user.Role == 1 {
		role = "seller"
	} else if user.Role == 2 {
		role = "buyer"
	} else {
		c.JSON(400, gin.H{"error": "Invalid role"})
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
		c.JSON(500, gin.H{"error": err.Error(), "success": false})
		return
	}

	user.ID = *res_id

	// Create a new token object, specifying signing method and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID":   user.ID,
		"username": user.Username,
		"role":     role,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte("secret"))
	if err != nil {
		c.JSON(500, gin.H{"error": "failed to generate token", "success": false})
		return
	}

	c.JSON(200, gin.H{"message": "User created successfully", "user": &user, "token": tokenString})
}

// ApproveSeller godoc
// @Method POST
// @Summary Approve seller
// @Description Approve seller by admin
// @Endpoint /api/v1/admin/approve-seller/:sellerID
func ApproveSeller(c *gin.Context, h *database.Handler) {
	sellerID := c.Param("sellerID")
	_, err := h.ChangeStatus(c, sellerID, "verified")
	if err != nil {
		fmt.Println("Error: ", err)
		c.JSON(500, gin.H{"success": false, "error": "failed to approve seller"})
		return
	}

	c.JSON(200, gin.H{"success": true, "data": bson.M{"seller_id": sellerID, "status": "verified"}})

}
