package services

import (
	"net/http"

	"github.com/2110366-2566-2/tortoise-app-backend/internal/database"
	"github.com/2110366-2566-2/tortoise-app-backend/internal/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TransactionHandler struct {
	handler *database.Handler
}

func NewTransactionHandler(handler *database.Handler) *TransactionHandler {
	return &TransactionHandler{handler: handler}
}

// GetTransactions godoc
// @Method GET
// @Summary Get transactions
// @Description Get transactions of user
// @Endpoint /api/v1/transactions/:userID
func (h *TransactionHandler) GetTransactions(c *gin.Context) {

	uid, _ := c.Get("userID")
	role, _ := c.Get("role")

	// Get transactions
	transactions, err := h.handler.GetTransactionByUserID(c, uid.(primitive.ObjectID), role.(string))

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	// Add details to the transaction
	for _, tx := range transactions {
		seller, _ := h.handler.GetUserByUserID(c, tx.SellerID.Hex())
		tx.SellerName = seller.FirstName + " " + seller.LastName

		buyer, _ := h.handler.GetUserByUserID(c, tx.BuyerID.Hex())
		tx.BuyerName = buyer.FirstName + " " + buyer.LastName

		pet, _ := h.handler.GetPetByPetID(c, tx.PetID.Hex())
		pet_detail := models.PetDetail{Name: pet.Name, Age: pet.Age, Sex: pet.Sex, Species: pet.Species, Media: pet.Media}
		tx.PetDetail = pet_detail
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data": gin.H{
			"transactions": transactions,
			"role":         role,
		},
	})
}

// GetTransactionByTransactionID godoc
// @Method GET
// @Summary Get transaction by transaction id
// @Description Get transaction by transaction id
// @Endpoint /api/v1/transactions/detail/:transactionID
func (h *TransactionHandler) GetTransactionByTransactionID(c *gin.Context) {
	transaction, err := h.handler.GetTransactionByTransactionID(c, c.Param("transactionID"))
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"success": true, "data": transaction})
}
