package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Buyer struct {
	ID        primitive.ObjectID `json:"id" bson:"_id"`
	FirstName string             `json:"first_name" bson:"first_name"`
	LastName  string             `json:"last_name" bson:"last_name"`
}
