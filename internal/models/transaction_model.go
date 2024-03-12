package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// transaction model

type Transaction struct {
	ID            primitive.ObjectID `json:"id" bson:"_id"`
	Price         int64              `json:"price" bson:"price"`
	BuyerID       primitive.ObjectID `json:"buyer_id" bson:"buyer_id"`
	SellerID      primitive.ObjectID `json:"seller_id" bson:"seller_id"`
	PetID         primitive.ObjectID `json:"pet_id" bson:"pet_id"`
	Status        string             `json:"status" bson:"status"`
	PaymentID     string             `json:"payment_id" bson:"payment_id"`
	PaymentMethod string             `json:"payment_method" bson:"payment_method"`
	Timestamp     time.Time          `json:"timestamp" bson:"timestamp"`
}

type PaymentIntent struct {
	ID            string             `json:"payment_id"`
	TransactionID primitive.ObjectID `json:"transaction_id"`
	PaymentMethod string             `json:"payment_method"`
}
