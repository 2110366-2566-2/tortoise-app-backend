package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// pet model from pet collection
type Pet struct {
	ID             primitive.ObjectID `bson:"_id,omitempty"`
	Name           string             `bson:"name"`
	Age            int32              `bson:"age"`
	Price          int32              `bson:"price"`
	Is_sold        bool               `bson:"is_sold"`
	Description    string             `bson:"description"`
	Weight         int32              `bson:"weight"`
	Sex            string             `bson:"sex"`
	Species        string             `bson:"species"`
	Type           string             `bson:"type"`
	Behavior       string             `bson:"behavior"`
	Media          string             `bson:"media"`
	Medical_record struct {
		Medical_id  string `bson:"medical_id"`
		Date        string `bson:"date"`
		Description string `bson:"description"`
	} `bson:"medical_record"`
	Seller_id string `bson:"seller_id"`
}

// pet card model from pet collection
type PetCard struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Name        string             `bson:"name"`
	Price       int32              `bson:"price"`
	Type        string             `bson:"type"`
	Media       string             `bson:"media"`
	Seller_id   string             `bson:"seller_id"`
	Seller_name string
}
