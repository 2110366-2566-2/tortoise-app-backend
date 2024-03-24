package apiV1

import (
	"github.com/2110366-2566-2/tortoise-app-backend/internal/database"
	"github.com/2110366-2566-2/tortoise-app-backend/internal/services"
	"github.com/gin-gonic/gin"
)

func TransactionServices(r *gin.RouterGroup, h *database.Handler) {
	// Create a new transaction handler
	transactionHandler := services.NewTransactionHandler(h)

	r.GET("/", transactionHandler.GetTransactions)
	r.GET("/:transactionID", transactionHandler.GetTransactionByTransactionID)
}
