package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// Seller model from seller collection
type Seller struct {
	ID        primitive.ObjectID   `json:"id" bson:"_id"`
	FirstName string               `json:"first_name"`
	LastName  string               `json:"last_surname"`
	Pets      []primitive.ObjectID `json:"pets"`
}
