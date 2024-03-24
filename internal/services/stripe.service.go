package services

import (
	"fmt"
	"log"

	"github.com/2110366-2566-2/tortoise-app-backend/internal/models"
	"github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/paymentintent"
)

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

func (h *PaymentHandler) CreatePaymentIntent(transaction models.Transaction) (*stripe.PaymentIntent, error) {
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
		return nil, fmt.Errorf("failed to create transaction")
	}

	return pi, nil
}

func (h *PaymentHandler) ConfirmPaymentIntent(transaction *models.Transaction, payment *models.PaymentIntent) error {
	// Set Stripe API key
	stripe.Key = h.env.STRIPE_KEY

	var PaymentMethod string

	if transaction.PaymentMethod == "promptpay" {
		// PaymentMethod = "promptpay"
		PaymentMethod = "pm_card_th_credit"
	} else {
		PaymentMethod = "pm_card_th_credit"
	}

	_, err := paymentintent.Confirm(
		payment.ID,
		&stripe.PaymentIntentConfirmParams{
			PaymentMethod: stripe.String(PaymentMethod),
		},
	)

	if err != nil {
		logStripeError(err)
		return fmt.Errorf("failed to confirm transaction")
	}

	return nil

}
