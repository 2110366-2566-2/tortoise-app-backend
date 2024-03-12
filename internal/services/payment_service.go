package services

import (
	"log"
	"time"

	"github.com/2110366-2566-2/tortoise-app-backend/configs"
	"github.com/2110366-2566-2/tortoise-app-backend/internal/database"
	"github.com/2110366-2566-2/tortoise-app-backend/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/paymentintent"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	var payment models.PaymentIntent

	c.BindJSON(&payment)

	petID, err := primitive.ObjectIDFromHex(payment.PetID)
	if err != nil {
		c.JSON(400, gin.H{"success": false, "error": "failed to convert pet ID"})
		return
	}

	isSold, err := h.handler.CheckPetStatus(c, petID)
	if err != nil {
		c.JSON(400, gin.H{"success": false, "error": "failed to check pet status"})
		return
	}

	if isSold {
		c.JSON(400, gin.H{"success": false, "error": "pet is not available"})
		return
	}

	// Set Stripe API key
	stripe.Key = h.env.STRIPE_KEY

	params := &stripe.PaymentIntentParams{
		Amount:             stripe.Int64(payment.Price * 100),
		Currency:           stripe.String(string(stripe.CurrencyTHB)),
		PaymentMethodTypes: stripe.StringSlice([]string{"card", "promptpay"}),
	}

	// Add metadata
	params.AddMetadata("pet_id", payment.PetID)
	params.AddMetadata("seller_id", payment.SellerID)
	params.AddMetadata("buyer_id", payment.BuyerID)

	pi, err := paymentintent.New(params)
	if err != nil {
		logStripeError(err)
		c.JSON(400, gin.H{"success": false, "error": "failed to create payment"})
		return
	}

	// convert petID to primitive.ObjectID
	petID, err1 := primitive.ObjectIDFromHex(payment.PetID)

	// Update pet status
	_, err2 := h.handler.UpdatePetStatus(c, petID, true)

	// Check if there is an error
	if err1 != nil || err2 != nil {
		// cancel payment intent
		_, err := paymentintent.Cancel(pi.ID, nil)
		if err != nil {
			c.JSON(400, gin.H{"success": false, "error": "failed to cancel payment"})
			return
		}
		c.JSON(400, gin.H{"success": false, "error": "failed to create payment"})
		return
	}

	// Add payment ID to payment model
	payment.ID = pi.ID

	c.JSON(200, gin.H{"success": true, "data": payment})

}

// ConfirmPayment godoc
// @Method POST
// @Summary Confirm payment
// @Description Confirm payment for pet
// @Endpoint /api/v1/payment/confirm
func (h *PaymentHandler) ConfirmPayment(c *gin.Context) {
	var transaction models.Transaction

	c.BindJSON(&transaction)

	// Set Stripe API key
	stripe.Key = h.env.STRIPE_KEY

	_, err := paymentintent.Confirm(
		transaction.PaymentID,
		&stripe.PaymentIntentConfirmParams{
			PaymentMethod: stripe.String("pm_card_visa"),
		},
	)

	// Check if there is an error
	if err != nil {
		logStripeError(err)
		// rollback
		_, err := h.handler.UpdatePetStatus(c, transaction.PetID, false)
		if err != nil {
			c.JSON(400, gin.H{"success": false, "error": "failed to rollback"})
			return
		}
		c.JSON(400, gin.H{"success": false, "error": "failed to confirm payment"})
		return
	}

	// Add timestamp
	transaction.Timestamp = time.Now()
	// Add transaction to database
	res, err := h.handler.CreateTransaction(c, &transaction)
	if err != nil {
		// rollback
		_, err := h.handler.UpdatePetStatus(c, transaction.PetID, false)
		if err != nil {
			c.JSON(400, gin.H{"success": false, "error": "failed to rollback"})
			return
		}
		c.JSON(400, gin.H{"success": false, "error": "failed to create transaction"})
		return
	}

	// ensure update pet status
	_, err = h.handler.UpdatePetStatus(c, transaction.PetID, true)
	if err != nil {
		c.JSON(400, gin.H{"success": false, "error": "failed to update pet status"})
		return
	}

	c.JSON(200, gin.H{"success": true, "data": res})
}
