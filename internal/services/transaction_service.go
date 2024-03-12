package services

import (
	"github.com/2110366-2566-2/tortoise-app-backend/internal/database"
	"github.com/gin-gonic/gin"
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
func (h *TransactionHandler) GetAllTransactions(c *gin.Context) {
	transactions, err := h.handler.GetTransactionByUserID(c, c.Param("userID"))
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"success": true, "data": transactions, "count": len(*transactions)})
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
