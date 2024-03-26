package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// pet model from pet collection
type Pet struct {
	ID              primitive.ObjectID `json:"id" bson:"_id"`
	Name            string             `json:"name"`
	Age             int32              `json:"age"`
	Price           int32              `json:"price"`
	Is_sold         bool               `json:"is_sold"`
	Description     string             `json:"description"`
	Weight          float64            `json:"weight"`
	Sex             string             `json:"sex"`
	Category        string             `json:"category"`
	Species         string             `json:"species"`
	Behavior        string             `json:"behavior"`
	Media           string             `json:"media"`
	Medical_records []Medical_record   `json:"medical_records"`
	Seller_id       primitive.ObjectID `json:"seller_id,omitempty"`
}

type Medical_record struct {
	Medical_id   string `json:"medical_id" bson:"medical_id"`
	Medical_date string `json:"medical_date" bson:"medical_date"`
	Description  string `json:"description"`
}

// pet card model from pet collection
type PetCard struct {
	ID             primitive.ObjectID `json:"id" bson:"_id"`
	Name           string             `json:"name"`
	Price          int                `json:"price"`
	Category       string             `json:"category"`
	Species        string             `json:"species"`
	Media          string             `json:"media"`
	Seller_id      primitive.ObjectID `json:"seller_id"`
	Seller_name    string             `json:"seller_name"`
	Seller_surname string             `json:"seller_surname"`
}

type PetDetail struct {
	Name    string `json:"name"`
	Age     int32  `json:"age"`
	Sex     string `json:"sex"`
	Species string `json:"species"`
	Media   string `json:"media"`
}

type PetMedia struct {
	ID    string `json:"id"`
	Media string `json:"media"`
}
