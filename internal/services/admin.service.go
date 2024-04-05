package services

import (
	"fmt"

	"github.com/2110366-2566-2/tortoise-app-backend/internal/database"
	"github.com/2110366-2566-2/tortoise-app-backend/internal/models"
	"github.com/2110366-2566-2/tortoise-app-backend/internal/storage"
	"github.com/2110366-2566-2/tortoise-app-backend/pkg/utils"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// AdminRegisterHandler godoc
// @Method POST
// @Summary Register admin
// @Description Register admin
// @Endpoint /api/v1/admin/register
func AdminRegisterHandler(c *gin.Context, h *database.Handler, storage *storage.Handler) {
	var admin models.Admin

	// Bind the request body to user model
	if err := c.BindJSON(&admin); err != nil {
		c.JSON(400, gin.H{"success": false, "error": err.Error()})
		return
	}

	// Hash the password
	admin.Password = utils.HashPassword(admin.Password)

	// Check if the user already exists
	if exist, err := database.CheckAdminExist(c, h, admin); err != nil {
		c.JSON(500, gin.H{"success": false, "error": err.Error()})
		return
	} else if exist {
		c.JSON(400, gin.H{"success": false, "error": "Username already exists"})
		return
	}

	// Check if the email already exists
	if exist, err := database.CheckAdminEmailExist(c, h, admin); err != nil {
		c.JSON(500, gin.H{"success": false, "error": "Internal server error"})
		return
	} else if exist {
		c.JSON(400, gin.H{"success": false, "error": "Email already exists"})
		return
	}

	admin.ID = primitive.NewObjectID()

	// convert image from base64 to url
	if admin.Image != "" {
		imageURL, err := storage.AddImage(c, admin.ID.Hex(), "admins", admin.Image)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error(), "success": false})
			return
		}
		admin.Image = imageURL
	}

	err := database.CreateAdmin(c, h, admin)
	// Create the user
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error(), "success": false})
		return
	}

	c.JSON(200, gin.H{"message": "Admin created successfully", "admin": &admin})
}

// LoginHandlerForAdmin godoc
// @Method POST
// @Summary Login admin
// @Description Login admin
// @Endpoint /api/v1/admin/login
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

	tokenString, err := utils.CreateTokenString(userID, username, role)
	if err != nil {
		fmt.Println("Error: ", err)
		c.JSON(500, gin.H{"error": "failed to generate token"})
		return
	}

	c.JSON(200, gin.H{"success": true, "token": tokenString})
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
