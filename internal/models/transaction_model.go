package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// transaction model

type Transaction struct {
	ID            primitive.ObjectID `json:"id" bson:"_id"`
	Price         int64              `json:"price"`
	BuyerID       primitive.ObjectID `json:"buyer_id"`
	SellerID      primitive.ObjectID `json:"seller_id"`
	PetID         primitive.ObjectID `json:"pet_id"`
	PaymentID     string             `json:"payment_id"`
	PaymentMethod string             `json:"payment_method"`
	Timestamp     time.Time          `json:"timestamp"`
}

type PaymentIntent struct {
	ID       string `json:"payment_id"`
	Price    int64  `json:"price"`
	BuyerID  string `json:"buyer_id"`
	SellerID string `json:"seller_id"`
	PetID    string `json:"pet_id"`
}
