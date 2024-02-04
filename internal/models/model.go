package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// user model from user collection
type User struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	Name         string             `bson:"name"`
	Surname      string             `bson:"surname"`
	Gender       string             `bson:"gender"`
	Phone_number string             `bson:"phone_number"`
	Image        string             `bson:"image"`
	Role         int32              `bson:"role"`
	Email        string             `bson:"email"`
	Password     string             `bson:"password"`
	Address      struct {
		Province     string `bson:"province"`
		District     string `bson:"district"`
		Sub_district string `bson:"sub_district"`
		Postal_code  string `bson:"postal_code"`
		Street       string `bson:"street"`
		Building     string `bson:"building"`
		House_number string `bson:"house_number"`
	} `bson:"address"`
	Pets []primitive.ObjectID `bson:"pets"`
}

// pet model pet collection
type Pet struct {
	ID             primitive.ObjectID `bson:"_id,omitempty"`
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
