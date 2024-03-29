package services

import (
	"fmt"
	"log"
	"time"

	"github.com/2110366-2566-2/tortoise-app-backend/configs"
	"github.com/2110366-2566-2/tortoise-app-backend/internal/database"
	"github.com/2110366-2566-2/tortoise-app-backend/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/stripe/stripe-go/v72"
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

func logStripeError(err error) {
	if stripeErr, ok := err.(*stripe.Error); ok {
		switch stripeErr.Type {
		case stripe.ErrorTypeCard:
			log.Println("A payment error occurred:", stripeErr.Msg)
		case stripe.ErrorTypeInvalidRequest:
			log.Println("An invalid request occurred.")
		default:
			log.Println("Another Stripe error occurred.")
		}
	} else {
		log.Println("An error occurred that was unrelated to Stripe.")
	}
}

// CreatePayment godoc
// @Method POST
// @Summary Create payment
// @Description Create payment for pet
// @Endpoint /api/v1/payment/create
func (h *PaymentHandler) CreatePayment(c *gin.Context) {
	var transaction models.Transaction

	if err := c.ShouldBindJSON(&transaction); err != nil {
		log.Println("failed to bind JSON: ", err)
		c.JSON(400, gin.H{"success": false, "error": "failed to bind JSON"})
		return // Return here
	}

	isSold, err := h.handler.CheckPetStatus(c, transaction.PetID)
	if err != nil {
		log.Println("failed to check pet status: ", err)
		c.JSON(400, gin.H{"success": false, "error": "failed to check pet status"})
		return // Return here
	}

	if isSold {
		log.Println("pet is not available")
		c.JSON(400, gin.H{"success": false, "error": "pet is not available"})
		return // Return here
	}

	// Set Stripe API key
	stripe.Key = h.env.STRIPE_KEY

	params := &stripe.PaymentIntentParams{
		Amount:             stripe.Int64(transaction.Price * 100),
		Currency:           stripe.String(string(stripe.CurrencyTHB)),
		PaymentMethodTypes: stripe.StringSlice([]string{"card", "promptpay"}),
	}

	// Add metadata
	params.AddMetadata("pet_id", transaction.PetID.Hex())
	params.AddMetadata("seller_id", transaction.SellerID.Hex())
	params.AddMetadata("buyer_id", transaction.BuyerID.Hex())

	pi, err := paymentintent.New(params)
	if err != nil {
		logStripeError(err)
		// log.Println("failed to create transaction: ", err)
		c.JSON(400, gin.H{"success": false, "error": "failed to create transaction"})
		return // Return here
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
			return // Return here
		}
		c.JSON(400, gin.H{"success": false, "error": "failed to create transaction"})
		return // Return here
	}

	// Add transaction ID to transaction model
	transaction.PaymentID = pi.ID

	// Add timestamp
	transaction.Timestamp = time.Now()

	// Add transaction status
	transaction.Status = "pending"

	// Add transaction to database
	_, err = h.handler.CreateTransaction(c, &transaction)

	if err != nil {
		fmt.Println("failed to create transaction: ", err)
		// rollback
		_, err1 := h.handler.UpdatePetStatus(c, transaction.PetID, false)
		if err1 != nil {
			c.JSON(400, gin.H{"success": false, "error": "failed to rollback"})
			return // Return here
		}
		// cancel transaction intent
		_, err2 := paymentintent.Cancel(pi.ID, nil)
		if err2 != nil {
			c.JSON(400, gin.H{"success": false, "error": "failed to cancel transaction"})
			return // Return here
		}
		c.JSON(400, gin.H{"success": false, "error": "failed to create transaction"})
		return // Return here
	}

	res := bson.M{
		"transaction_id": transaction.ID.Hex(),
		"payment_id":     pi.ID,
	}

	c.JSON(200, gin.H{"success": true, "data": res})
}

func (h *PaymentHandler) AA(c *gin.Context) {
	// Set Stripe API key
}

// ConfirmPayment godoc
// @Method POST
// @Summary Confirm payment
// @Description Confirm payment for pet
// @Endpoint /api/v1/payment/confirm
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
		c.JSON(400, gin.H{"success": false, "error": "failed to update transaction"})
		return // Return here
	}

	var transaction models.Transaction

	err = res.Decode(&transaction)
	if err != nil {
		c.JSON(400, gin.H{"success": false, "error": "failed to decode transaction"})
		return
	}

	// Set Stripe API key
	stripe.Key = h.env.STRIPE_KEY

	var PaymentMethod string

	if transaction.PaymentMethod == "promptpay" {
		// PaymentMethod = "promptpay"
		PaymentMethod = "pm_card_th_credit"
	} else {
		PaymentMethod = "pm_card_th_credit"
	}

	_, err = paymentintent.Confirm(
		payment.ID,
		&stripe.PaymentIntentConfirmParams{
			PaymentMethod: stripe.String(PaymentMethod),
		},
	)

	// Check if there is an error
	if err != nil {
		logStripeError(err)
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
