package services

import (
	"github.com/2110366-2566-2/tortoise-app-backend/internal/database"
	// "github.com/2110366-2566-2/tortoise-app-backend/internal/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"github.com/gin-gonic/gin"
	// "github.com/dgrijalva/jwt-go"
	//"time"
	"net/http"
	// "fmt"
)

func GetTransactions(c *gin.Context, h *database.Handler) {
	uid, _ := c.Get("userID")
	role, _ := c.Get("role")

	transactions, err := database.GetTransactionByID(c, h, uid.(primitive.ObjectID), role.(string))

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, transactions)
}