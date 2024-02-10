package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// pet model from pet collection
type Pet struct {
	ID             primitive.ObjectID `json:"id" bson:"_id"`
	Name           string             `json:"name"`
	Age            int32              `json:"age"`
	Price          int32              `json:"price"`
	Is_sold        bool               `json:"is_sold"`
	Description    string             `json:"description"`
	Weight         float64            `json:"weight"`
	Sex            string             `json:"sex"`
	Species        string             `json:"species"`
	Type           string             `json:"type"`
	Behavior       string             `json:"behavior"`
	Media          string             `json:"media"`
	Medical_record struct {
		Medical_id  string `json:"medical_id"`
		Date        string `json:"date"`
		Description string `json:"description"`
	} `json:"medical_record"`
	Seller_id primitive.ObjectID `json:"seller_id,omitempty"`
}

// pet card model from pet collection
type PetCard struct {
	ID             primitive.ObjectID `json:"id" bson:"_id"`
	Name           string             `json:"name"`
	Price          int                `json:"price"`
	Type           string             `json:"type"`
	Species        string             `json:"species"`
	Media          string             `json:"media"`
	Seller_id      primitive.ObjectID `json:"seller_id"`
	Seller_name    string             `json:"seller_name"`
	Seller_surname string             `json:"seller_surname"`
}
