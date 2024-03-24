package apiV1

import (
	"github.com/2110366-2566-2/tortoise-app-backend/internal/database"
	"github.com/2110366-2566-2/tortoise-app-backend/internal/services"
	"github.com/gin-gonic/gin"
)

func BankServices(r *gin.RouterGroup, h *database.Handler) {
	// Create a new seller handler
	sellerHandler := services.NewSellerHandler(h)

	// Set up routes
	r.POST("/:sellerID", sellerHandler.AddBankAccount)
	r.GET("/:sellerID", sellerHandler.GetBankAccount)
	r.DELETE("/:sellerID", sellerHandler.DeleteBankAccount)
}
