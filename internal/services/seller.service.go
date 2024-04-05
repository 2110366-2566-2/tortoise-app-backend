package services

import (
	"fmt"
	"strings"

	"github.com/2110366-2566-2/tortoise-app-backend/internal/database"
	"github.com/2110366-2566-2/tortoise-app-backend/internal/models"
	"github.com/gin-gonic/gin"
)

type SellerHandler struct {
	handler *database.Handler
}

func NewSellerHandler(handler *database.Handler) *SellerHandler {
	return &SellerHandler{handler: handler}
}

// AddBankAccount godoc
// @Method POST
// @Summary Add bank account
// @Description Add bank account to seller
// @Endpoint /api/v1/bank/:sellerID
func (h *SellerHandler) AddBankAccount(c *gin.Context) {
	var bankAccount models.BankAccount
	if err := c.ShouldBindJSON(&bankAccount); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	res, err := h.handler.AddBankAccount(c, c.Param("sellerID"), bankAccount)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, gin.H{"success": true, "result": res})
}

// GetBankAccount godoc
// @Method GET
// @Summary Get bank account
// @Description Get bank account of seller
// @Endpoint /api/v1/bank/:sellerID
func (h *SellerHandler) GetBankAccount(c *gin.Context) {
	bankAccount, err := h.handler.GetBankAccount(c.Param("sellerID"))
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"success": true, "data": bankAccount})
}

// DeleteBankAccount godoc
// @Method DELETE
// @Summary Delete bank account
// @Description Delete bank account of seller
// @Endpoint /api/v1/bank/:sellerID
func (h *SellerHandler) DeleteBankAccount(c *gin.Context) {
	res, err := h.handler.DeleteBankAccount(c.Param("sellerID"))
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"success": true, "data": res})
}

// GetSeller godoc
// @Method GET
// @Summary Get seller by ID
// @Description Get seller by ID
// @Endpoint /api/v1/seller/:sellerID
func (h *SellerHandler) GetSeller(c *gin.Context) {
	seller, err := h.handler.GetSellerBySellerID(c, c.Param("sellerID"))
	if err != nil {
		fmt.Println("Error: ", err)
		if strings.Contains(err.Error(), "failed to find seller") {
			c.JSON(500, gin.H{"success": false, "error": "seller not found"})
			return
		}
		c.JSON(500, gin.H{"error": "failed to get seller", "success": false})
		return
	}

	c.JSON(200, gin.H{"success": true, "data": seller})
}

// GetAllSellers godoc
// @Method GET
// @Summary Get all sellers with query params
// @Description Get all sellers
// @Endpoint /api/v1/sellers
func (h *SellerHandler) GetAllSellers(c *gin.Context) {
	// query
	status := c.Query("status")
	sellers, err := h.handler.GetAllSellers(c, status)
	if err != nil {
		fmt.Println("Error: ", err)
		if strings.Contains(err.Error(), "invalid status") {
			c.JSON(400, gin.H{"success": false, "error": "invalid status"})
			return
		}
		c.JSON(500, gin.H{"error": "failed to get sellers", "success": false})
		return
	}

	c.JSON(200, gin.H{"success": true, "data": sellers, "count": len(*sellers)})
}
