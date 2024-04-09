package services

import (
	"fmt"
	"log"
	"time"

	"github.com/2110366-2566-2/tortoise-app-backend/configs"
	"github.com/2110366-2566-2/tortoise-app-backend/internal/database"
	"github.com/2110366-2566-2/tortoise-app-backend/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/stripe/stripe-go/v72/paymentintent"
	"go.mongodb.org/mongo-driver/bson"
)

type PaymentHandler struct {
	handler *database.Handler
	env     configs.EnvVars
}

func NewPaymentHandler(handler *database.Handler, env configs.EnvVars) *PaymentHandler {
	return &PaymentHandler{
		handler: handler,
		env:     env,
	}
}

// CreatePayment godoc
// @Summary Create a new payment
// @Description Create a new payment for a transaction
// @Security ApiKeyAuth
// @Tags Payments
// @Accept json
// @Produce json
// @Router /payment/create [post]
// @Param payment body models.CreatePaymentBody true "Payment object"
// @Success 201 {object} models.CreatePaymentResponse
// @Failure 400 {object} models.ErrorResponse
func (h *PaymentHandler) CreatePayment(c *gin.Context) {
	var transaction models.Transaction

	if err := c.ShouldBindJSON(&transaction); err != nil {
		log.Println("failed to bind JSON: ", err)
		c.JSON(400, gin.H{"success": false, "error": "failed to bind JSON"})
		return
	}

	isSold, err := h.handler.CheckPetStatus(c, transaction.PetID)
	if err != nil {
		log.Println("failed to check pet status: ", err)
		c.JSON(400, gin.H{"success": false, "error": "failed to check pet status"})
		return
	}

	if isSold {
		log.Println("pet is not available")
		c.JSON(400, gin.H{"success": false, "error": "pet is not available"})
		return
	}

	pi, err := h.CreatePaymentIntent(transaction)
	if err != nil {
		// log.Println("failed to create transaction: ", err)
		c.JSON(400, gin.H{"success": false, "error": "failed to create transaction"})
		return
	}

	// Update pet status
	_, err = h.handler.UpdatePetStatus(c, transaction.PetID, true)

	// Check if there is an error
	if err != nil {
		// cancel transaction intent
		log.Println("failed to cancel transaction")
		_, err := paymentintent.Cancel(pi.ID, nil)
		if err != nil {
			c.JSON(400, gin.H{"success": false, "error": "failed to cancel transaction"})
			return
		}
		c.JSON(400, gin.H{"success": false, "error": "failed to create transaction"})
		return
	}

	// Add transaction ID, payment status, and timestamp
	transaction.PaymentID = pi.ID
	transaction.Status = "pending"
	transaction.Timestamp = time.Now()

	// Add transaction to database
	_, err = h.handler.CreateTransaction(c, &transaction)

	if err != nil {
		fmt.Println("failed to create transaction: ", err)
		// rollback
		_, err1 := h.handler.UpdatePetStatus(c, transaction.PetID, false)
		if err1 != nil {
			fmt.Println("failed to rollback: ", err1)
			c.JSON(400, gin.H{"success": false, "error": "failed to create transaction"})
			return
		}
		// cancel transaction intent
		_, err2 := paymentintent.Cancel(pi.ID, nil)
		if err2 != nil {
			fmt.Println("failed to cancel transaction: ", err2)
			c.JSON(400, gin.H{"success": false, "error": "failed to create transaction"})
			return
		}
		fmt.Println("failed to create transaction: ", err)
		c.JSON(400, gin.H{"success": false, "error": "failed to create transaction"})
		return
	}

	c.JSON(201, gin.H{"success": true, "data": bson.M{
		"transaction_id": transaction.ID.Hex(),
		"payment_id":     pi.ID,
	}})
}

// ConfirmPayment godoc
// @Summary Confirm a payment
// @Description Confirm a payment for a transaction
// @Security ApiKeyAuth
// @Tags Payments
// @Accept json
// @Produce json
// @Param Payment body models.PaymentIntent true "Payment object"
// @Router /payment/confirm [post]
// @Success 200 {object} models.ConfirmPaymentResponse
// @Failure 400 {object} models.ErrorResponse
func (h *PaymentHandler) ConfirmPayment(c *gin.Context) {
	var payment models.PaymentIntent

	c.BindJSON(&payment)

	update := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "status", Value: "paid"},
			{Key: "payment_method", Value: payment.PaymentMethod},
		}},
	}

	// Update transaction
	res, err := h.handler.UpdateTransaction(c, payment.TransactionID, update)

	if err != nil {
		fmt.Println("failed to update transaction")
		c.JSON(400, gin.H{"success": false, "error": "failed to update transaction"})
		return // Return here
	}

	var transaction models.Transaction

	err = res.Decode(&transaction)
	if err != nil {
		c.JSON(400, gin.H{"success": false, "error": "failed to decode transaction"})
		return
	}

	err = h.ConfirmPaymentIntent(&transaction, &payment)

	// Check if there is an error
	if err != nil {
		// rollback
		_, err = h.handler.UpdateTransaction(c, payment.TransactionID, bson.D{{Key: "$set", Value: bson.D{{Key: "status", Value: "failed"}}}})
		if err != nil {
			c.JSON(400, gin.H{"success": false, "error": "failed to rollback"})
			return // Return here
		}
		_, err := h.handler.UpdatePetStatus(c, transaction.PetID, false)
		if err != nil {
			c.JSON(400, gin.H{"success": false, "error": "failed to rollback"})
			return // Return here
		}
		c.JSON(400, gin.H{"success": false, "error": "failed to confirm payment"})
		return // Return here
	}

	// ensure sold status is true
	_, err = h.handler.UpdatePetStatus(c, transaction.PetID, true)
	if err != nil {
		c.JSON(400, gin.H{"success": false, "error": "failed to update pet status"})
		return // Return here
	}

	c.JSON(200, gin.H{"success": true, "data": map[string]interface{}{"transaction_id": transaction.ID.Hex()}})
}
