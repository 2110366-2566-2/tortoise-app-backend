package models

import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Transaction struct {
	ID            primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	SellerID      primitive.ObjectID `json:"seller_id" bson:"seller_id" binding:"required"`
	PetID         primitive.ObjectID `json:"pet_id" bson:"pet_id" binding:"required"`
	BuyerID       primitive.ObjectID `json:"buyer_id" bson:"buyer_id" binding:"required"`
	Price         int64            	 `json:"price" binding:"required"`
	Timestamp     time.Time          `json:"timestamp" binding:"required"`
	PaymentMethod string             `json:"payment_method" bson:"payment_method" binding:"required"`
	Status 		  string 			 `json:"status" binding:"required"`
}

type TransactionWithDetails struct {
	Transaction	Transaction
	SellerName string `json:"seller_name"`
	BuyerName string `json:"buyer_name"`
	PetName string `json:"pet_name"`
}