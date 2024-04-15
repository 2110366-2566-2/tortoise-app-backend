package services

import (
	"log"
	"net/http"

	"github.com/2110366-2566-2/tortoise-app-backend/internal/database"
	"github.com/2110366-2566-2/tortoise-app-backend/internal/models"
	"github.com/2110366-2566-2/tortoise-app-backend/internal/storage"
	"github.com/2110366-2566-2/tortoise-app-backend/pkg/utils"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserHandler struct {
	dbHandler  *database.Handler
	stgHandler *storage.Handler
}

func NewUserHandler(db *database.Handler, stg *storage.Handler) *UserHandler {
	return &UserHandler{
		dbHandler:  db,
		stgHandler: stg,
	}
}

func (h *UserHandler) GetUserByUserID(c *gin.Context) {
	id := c.Param("userID")
	user, err := h.dbHandler.GetUserByUserID(c, id)
	if err != nil {
		log.Println("Error: ", err)
		errorMsg := "failed to get user by user id"
		if err.Error() == "mongo: no documents in result" {
			errorMsg = "user not found"
		}
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": errorMsg})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "data": &user})
}

func (h *UserHandler) UpdateUserPasswd(c *gin.Context) {
	var data bson.M
	c.BindJSON(&data)

	//check json body
	if len(data) != 2 {
		c.JSON(400, gin.H{"success": false, "error": "less or more than two fields in request body"})
		return
	}

	for k, v := range data {
		if k != "password" && k != "old_password" {
			c.JSON(400, gin.H{"success": false, "error": "field name is incorrect"})
			return
		}

		if k == "old_password" {
			id := c.Param("userID")
			user, err := h.dbHandler.GetUserByUserID(c, id)

			if err != nil {
				log.Println("Error: ", err)
				errorMsg := "failed to get user by user id"
				if err.Error() == "mongo: no documents in result" {
					errorMsg = "user not found"
				}
				c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": errorMsg})
				return
			}

			if !utils.ComparePassword(user.Password, v.(string)) {
				c.JSON(http.StatusForbidden, gin.H{"success": false, "error": "password is incorrect"})
				return
			}
		}
	}

	//call funtion
	res, err := h.dbHandler.UpdateOneUser(c, c.Param("userID"), data)
	if err != nil {
		log.Println("Error: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": "failed to update user's password"})
		return
	}
	// log result
	var updatedUser models.User
	err = res.Decode(&updatedUser)
	if err != nil {
		log.Println("Error decoding updated user's password:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": "failed to decode updated user's password"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "data": updatedUser})
}

func (h *UserHandler) UpdateUser(c *gin.Context) {
	var data bson.M
	c.BindJSON(&data)

	//Check if body have "password field"
	if _, ok := data["password"]; ok {
		c.JSON(400, gin.H{"success": false, "error": "found password field in body"})
		return
	}

	// log data
	log.Println("Length of data: ", len(data))
	for k, v := range data {
		log.Println("Key: ", k, "Value: ", v)
	}

	utils.BsonSanitize(&data)

	// check if have media to upload
	if image, ok := data["image"]; ok {
		log.Println("Found image in body")
		url, err := h.stgHandler.AddImage(c, c.Param("userID"), "users", image.(string))
		if err != nil {
			if err.Error() == "invalid base64 image string" {
				c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "invalid base64 image string"})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": "failed to upload media"})
			return
		}
		data["image"] = url
	}

	res, err := h.dbHandler.UpdateOneUser(c, c.Param("userID"), data)
	if err != nil {
		log.Println("Error: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": "failed to update user's profiles"})
		return
	}
	// log result
	var updatedUser models.User
	err = res.Decode(&updatedUser)
	if err != nil {
		log.Println("Error decoding updated user's password:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": "failed to decode updated user's profiles"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "data": updatedUser})
}

func (h *UserHandler) DeleteUser(c *gin.Context) {
	userID := utils.SanitizeString(c.Param("userID"))
	user, err := h.dbHandler.GetUserByUserID(c, userID)
	if err != nil {
		log.Println("Error: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": "user not found"})
		return
	}

	// delete user
	res, err := h.dbHandler.DeleteOneUser(c, userID, h.stgHandler)
	if err != nil {
		log.Println("Error: ", err)
		errorMsg := "failed to delete user"
		if err.Error()[:18] == "failed to get user:" {
			errorMsg = "user not found"
		}
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": errorMsg})
		return
	}

	// if user have image, delete it
	if len(user.Image) > 0 {
		// delete media
		if err := h.stgHandler.DeleteImage(c, userID, "users"); err != nil {
			log.Println("Error: ", err)
			c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": "failed to delete media"})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "deletedCount": res.DeletedCount})
}

func (h *UserHandler) UpdateForgotPassword(c *gin.Context) {
	var reset models.ResetPassword
	if err := c.BindJSON(&reset); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "failed to bind data"})
		return
	}

	var user *models.User
	user, err := h.dbHandler.GetUserByMail(c, bson.M{"email": reset.Email})
	if err != nil {
		log.Println("Error: ", err)
		errorMsg := "failed to get user by email"
		if err.Error() == "mongo: no documents in result" {
			errorMsg = "user not found"
		}
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": errorMsg})
		return
	}

	// validate user_id with token
	IDFromToken, exist := c.Get("userID")
	if !exist {
		c.JSON(http.StatusUnauthorized, gin.H{"success": false, "error": "Unauthorized"})
		return
	}
	if IDFromToken.(primitive.ObjectID) != user.ID {
		c.JSON(http.StatusUnauthorized, gin.H{"success": false, "error": "Unauthorized"})
		return
	}

	// Update user's password
	update := bson.M{
		"password": reset.Password,
	}
	_, err = h.dbHandler.UpdateOneUser(c, user.ID.Hex(), update)
	if err != nil {
		log.Println("Error: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": "failed to update user's password"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "data": "password updated"})

}

func (h *UserHandler) WhoAmI(c *gin.Context) {
	userID, exist := c.Get("userID")
	if !exist {
		c.JSON(http.StatusUnauthorized, gin.H{"success": false, "error": "Unauthorized"})
		return
	}
	// convert primitive.ObjectID to string
	userIDStr := userID.(primitive.ObjectID).Hex()

	// fmt.Println("userID: ", userIDStr)
	user, err := h.dbHandler.GetUserByUserID(c, userIDStr)
	if err != nil {
		log.Println("Error: ", err)
		errorMsg := "failed to get user by user id"
		if err.Error() == "failed to find user" {
			errorMsg = "user not found"
		}
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": errorMsg})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "data": user})
}

func (h *UserHandler) ValidatePassword(c *gin.Context) {

	var data models.Password
	c.BindJSON(&data)

	// Prevent XSS attack
	password := utils.SanitizeString(data.Password)

	// log.Println("Password:", password)

	if len(password) <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "incorrect password format"})
		return
	}

	userID, exist1 := c.Get("userID")
	role, exist2 := c.Get("role")
	if !exist1 || !exist2 {
		c.JSON(http.StatusUnauthorized, gin.H{"success": false, "error": "Unauthorized"})
		return
	}

	var hashedPassword string

	// log.Println("Role:", role)

	// convert primitive.ObjectID to string
	if role == "admin" {
		// get admin by user id
		admin, err := database.GetAdminByUserID(c, h.dbHandler, userID.(primitive.ObjectID))
		if err != nil {
			log.Println("Error: ", err)
			errorMsg := "failed to validate password"
			c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": errorMsg})
			return
		}
		hashedPassword = admin.Password
	} else {
		userID = userID.(primitive.ObjectID).Hex()
		// get user by user id
		user, err := h.dbHandler.GetUserByUserID(c, userID.(string))
		if err != nil {
			log.Println("Error: ", err)
			errorMsg := "failed to validate password"
			c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": errorMsg})
			return
		}
		hashedPassword = user.Password
	}
	// validate password
	if !utils.ComparePassword(hashedPassword, password) {
		c.JSON(http.StatusForbidden, gin.H{"success": false, "message": "password is incorrect"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "message": "password is correct"})

}
