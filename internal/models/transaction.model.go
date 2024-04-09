package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// transaction model
type Transaction struct {
	ID            primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	SellerID      primitive.ObjectID `json:"seller_id" bson:"seller_id" binding:"required"`
	PetID         primitive.ObjectID `json:"pet_id" bson:"pet_id" binding:"required"`
	BuyerID       primitive.ObjectID `json:"buyer_id" bson:"buyer_id" binding:"required"`
	PaymentID     string             `json:"payment_id" bson:"payment_id"`
	Price         int64              `json:"price" binding:"required"`
	PaymentMethod string             `json:"payment_method" bson:"payment_method"`
	Status        string             `json:"status"`
	Timestamp     time.Time          `json:"timestamp"`
}

type PaymentIntent struct {
	ID            string             `json:"payment_id" bson:"payment_id" example:"123456789"`
	TransactionID primitive.ObjectID `json:"transaction_id" bson:"transaction_id" example:"60163b3be1e8712c4e7f35cf"`
	PaymentMethod string             `json:"payment_method" bson:"payment_method" example:"card" enum:"card,promptpay"`
}

type TransactionWithDetails struct {
	Transaction
	SellerName string    `json:"seller_name" bson:"seller_name"`
	BuyerName  string    `json:"buyer_name" bson:"buyer_name"`
	PetDetail  PetDetail `json:"pet_detail" bson:"pet_detail"`
}
