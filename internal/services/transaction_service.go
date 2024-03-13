package services

import (
	"github.com/2110366-2566-2/tortoise-app-backend/internal/database"
	"github.com/2110366-2566-2/tortoise-app-backend/internal/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"net/http"
)

func GetTransactions(c *gin.Context, h *database.Handler) {

	uid, _ := c.Get("userID")
	role, _ := c.Get("role")

	transactions, err := database.GetTransactionByID(c, h, uid.(primitive.ObjectID), role.(string))

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	// Add details to the transaction
	for _, tx := range transactions {
		seller, _ := h.GetUserByUserID(c, tx.SellerID.Hex())
		tx.SellerName = seller.FirstName + " " + seller.LastName

		buyer, _ := h.GetUserByUserID(c, tx.BuyerID.Hex())
		tx.BuyerName = buyer.FirstName + " " + buyer.LastName

		pet, _ := h.GetPetByPetID(c, tx.PetID.Hex())
		pet_detail := models.PetDetail{Name: pet.Name, Age: pet.Age, Sex: pet.Sex, Species: pet.Species}
		tx.PetDetail = pet_detail
	}

	c.JSON(http.StatusOK, gin.H{"role": role, "data": transactions})
}
