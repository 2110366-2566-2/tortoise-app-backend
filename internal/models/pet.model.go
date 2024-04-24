package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// pet model from pet collection
type Pet struct {
	ID              primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty" example:"60163b3be1e8712c4e7f35cf"`
	Name            string             `json:"name" example:"Fluffy"`
	Age             float32            `json:"age" example:"3.0"`
	Price           int32              `json:"price" example:"500"`
	Is_sold         bool               `json:"is_sold" example:"false"`
	Description     string             `json:"description" example:"A friendly and playful dog"`
	Weight          float64            `json:"weight" example:"25.5"`
	Sex             string             `json:"sex" example:"male"`
	Category        string             `json:"category" example:"Dog"`
	Species         string             `json:"species" example:"Golden Retriever"`
	Behavior        string             `json:"behavior" example:"Friendly"`
	Media           string             `json:"media" example:"https://example.com/fluffy.jpg"`
	Medical_records []Medical_record   `json:"medical_records" bson:"medical_records"`
	Seller_id       primitive.ObjectID `json:"seller_id,omitempty" bson:"seller_id,omitempty" example:"60163b3be1e8712c4e7f35ce"`
}

type Medical_record struct {
	Medical_id   string `json:"medical_id" bson:"medical_id" example:"123456789"`
	Medical_date string `json:"medical_date" bson:"medical_date" example:"2024-04-08"`
	Description  string `json:"description" example:"Routine checkup"`
}

type PetCard struct {
	ID            primitive.ObjectID `json:"id" bson:"_id" example:"60163b3be1e8712c4e7f35cf"`
	Name          string             `json:"name" example:"Fluffy"`
	Price         int                `json:"price" example:"100"`
	Category      string             `json:"category" example:"Dog"`
	Species       string             `json:"species" example:"Golden Retriever"`
	Media         string             `json:"media" example:"https://example.com/fluffy.jpg"`
	SellerID      primitive.ObjectID `json:"seller_id" bson:"seller_id" example:"60163b3be1e8712c4e7f35ce"`
	SellerName    string             `json:"seller_name" example:"John"`
	SellerSurname string             `json:"seller_surname" example:"Doe"`
}

type PetDetail struct {
	Name    string  `json:"name"`
	Age     float32 `json:"age"`
	Sex     string  `json:"sex"`
	Species string  `json:"species"`
	Media   string  `json:"media"`
}
