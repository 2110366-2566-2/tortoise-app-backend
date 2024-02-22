package services

import (
	"fmt"
	"log"

	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/charge"
)

func Payment() {
	// Set Stripe API key
	stripe.Key = "sk_test_51OktQDLXsvklyMmPEA5ELIb0pLN0dSMBGHmlTp5GtrFc1sVkrZxcHNK3UmgLWLCxCYtYF5n3pXmZ9A4khYNnRMTR00iWcR4YX7"

	// Create a new charge
	params := &stripe.ChargeParams{
		Amount:   stripe.Int64(2000),
		Currency: stripe.String(string(stripe.CurrencyTHB)),
	}
	params.SetSource("tok_visa") // use a test card token provided by Stripe
	ch, err := charge.New(params)

	// Check for errors
	if err != nil {
		log.Fatal(err)
	}

	// Print charge details
	fmt.Printf("Charge ID: %s\n", ch.ID)
	fmt.Printf("Amount: %d\n", ch.Amount)
	fmt.Printf("Description: %s\n", ch.Description)
}
