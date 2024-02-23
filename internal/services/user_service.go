package services

import (
	"log"
	"net/http"

	"github.com/2110366-2566-2/tortoise-app-backend/internal/database"
	"github.com/2110366-2566-2/tortoise-app-backend/internal/models"
	"github.com/2110366-2566-2/tortoise-app-backend/pkg/utils"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

type UserHandler struct {
	handler *database.Handler
}

func NewUserHandler(handler *database.Handler) *UserHandler {
	return &UserHandler{handler: handler}
}

// GetUserByUserID godoc
// @Method GET
// @Summary Get user by user id
// @Description Get user by user id
// @Endpoint /api/v1/user/:userID
func (h *UserHandler) GetUserByUserID(c *gin.Context) {
	id := c.Param("userID")
	user, err := h.handler.GetUserByUserID(c, id)
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

// UpdatePasswd godoc
// @Method PUT
// @Summary Update user's password
// @Description Update user's password by user id
// @Endpoint /api/v1/user/passwd/:userID
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
			user, err := h.handler.GetUserByUserID(c, id)

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
	res, err := h.handler.UpdateOneUser(c, c.Param("userID"), data)
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

// UpdateUser godoc
// @Method PUT
// @Summary Update user's profile
// @Description Update user's profile
// @Endpoint /api/v1/user/:userID
func (h *UserHandler) UpdateUser(c *gin.Context) {
	var data bson.M
	c.BindJSON(&data)

	//Check if body have "password field"
	for k, _ := range data {
		if k == "password" {
			c.JSON(400, gin.H{"success": false, "error": "found password field in body"})
			return
		}
	}
	res, err := h.handler.UpdateOneUser(c, c.Param("userID"), data)
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

// DeleteUser godoc
// @Method DELETE
// @Summary Delete user
// @Description Delete user that appears in all schema and delete pet if user is a seller
// @Endpoint /api/v1/user/:userID
func (h *UserHandler) DeleteUser(c *gin.Context) {
	res, err := h.handler.DeleteOneUser(c, c.Param("userID"))
	if err != nil {
		log.Println("Error: ", err)
		errorMsg := "failed to delete user"
		if err.Error()[:18] == "failed to get user:" {
			errorMsg = "user not found"
		}
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": errorMsg})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "deletedCount": res.DeletedCount})
}
