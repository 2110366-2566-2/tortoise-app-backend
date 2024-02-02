package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// user model
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
}

// pet model
type Animal struct {
	ID             string `bson:"_id"`
	Age            string `bson:"age"`
	Price          string `bson:"price"`
	Is_sold        bool   `bson:"is_sold"`
	Description    string `bson:"description"`
	Weight         string `bson:"weight"`
	Sex            string `bson:"sex"`
	Species        string `bson:"species"`
	Type           string `bson:"type"`
	Behavior       string `bson:"behavior"`
	Media          string `bson:"media"`
	Seller_id      string `bson:"seller_id"`
	Medical_record struct {
		Medical_id  string    `bson:"medical_id"`
		Date        time.Time `bson:"date"`
		Description string    `bson:"description"`
	} `bson:"medical_record"`
}
