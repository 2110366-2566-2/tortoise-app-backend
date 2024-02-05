package services

import (
	"fmt"
	"net/http"

	"github.com/2110366-2566-2/tortoise-app-backend/internal/database"
	// "github.com/2110366-2566-2/tortoise-app-backend/internal/models"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	handler *database.Handler
}

func NewUserHandler(handler *database.Handler) *UserHandler {
	return &UserHandler{handler: handler}
}

// GetAllUsers godoc
// @Method GET
// @Summary Get all users
// @Description Get all users collection
// @Endpoint /api/v1/users
func (h *UserHandler) GetAllUsers(c *gin.Context) {
	users, err := h.handler.GetAllUsers(c)
	if err != nil {
		fmt.Println("Error: ", err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, &users)
}

// GetUserByUserID godoc
// @Method GET
// @Summary Get user by user id
// @Description Get user by user id
// @Endpoint /api/v1/users/:userID
func (h *UserHandler) GetUserByUserID(c *gin.Context) {
	id := c.Param("userID")
	user, err := h.handler.GetUserByUserID(c, id)
	if err != nil {
		fmt.Println("Error: ", err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, &user)
}
