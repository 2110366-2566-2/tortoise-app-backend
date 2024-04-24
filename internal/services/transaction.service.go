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

func (h *TransactionHandler) GetTransactions(c *gin.Context) {

	// fmt.Println(c)
	uid, exits1 := c.Get("userID")
	role, exits2 := c.Get("role")

	// type of uid is primitive.ObjectID
	// fmt.Println(uid)
	// fmt.Println("Type of uid: ", reflect.TypeOf(uid))
	// fmt.Println("Type of role: ", reflect.TypeOf(role))
	// fmt.Println("Role: ", role)
	// fmt.Println("Context: ", c)

	if !exits1 {
		c.JSON(http.StatusNotFound, gin.H{"error": "failed to get userID"})
		return
	}

	if !exits2 {
		c.JSON(http.StatusNotFound, gin.H{"error": "failed to get role"})
		return
	}

	// Get transactions
	transactions, err := h.handler.GetTransactionByID(c, uid.(primitive.ObjectID), role.(string))

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

func (h *TransactionHandler) GetTransactionByTransactionID(c *gin.Context) {
	transaction, err := h.handler.GetTransactionByTransactionID(c, c.Param("transactionID"))
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"success": true, "data": transaction})
}
