package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Seller struct {
	ID      primitive.ObjectID   `bson:"_id,omitempty" json:"_id,omitempty"`
	Name    string               `json:"name"`
	Surname string               `json:"surname"`
	Pets    []primitive.ObjectID `json:"pets"`
}
