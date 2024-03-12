package services

import (
	"log"

	"github.com/2110366-2566-2/tortoise-app-backend/configs"
	"github.com/2110366-2566-2/tortoise-app-backend/internal/database"
	"github.com/2110366-2566-2/tortoise-app-backend/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/charge"
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

func (h *PaymentHandler) CreatePayment(c *gin.Context) {
	var transaction models.Transaction

	transaction.ID = primitive.NewObjectID()

	// Set Stripe API key
	stripe.Key = h.env.STRIPE_KEY

	// Create a new charge
	params := &stripe.ChargeParams{
		Amount:      stripe.Int64(transaction.Price * 100),
		Currency:    stripe.String(string(stripe.CurrencyTHB)),
		Description: stripe.String(transaction.ID.Hex()),
	}

	params.SetSource("tok_visa") // mock visa card

	meta_data := map[string]string{
		"transaction_id": transaction.ID.Hex(),
		"buyer_id":       transaction.BuyerID,
		"seller_id":      transaction.SellerID,
		"pet_id":         transaction.PetID,
	}
	params.Metadata = meta_data

	_, err := charge.New(params)

	if err != nil {
		log.Println("Error: ", err)
		c.JSON(400, gin.H{"success": false, "error": "payment failed"})
		return
	}

	// Save the transaction to the database
}
