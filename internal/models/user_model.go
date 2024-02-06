package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// user model from user collection
type User struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Name         string             `json:"name"`
	Surname      string             `json:"surname"`
	Gender       string             `json:"gender"`
	Phone_number string             `json:"phone_number"`
	Image        string             `json:"image"`
	Role         int32              `json:"role"`
	Email        string             `json:"email"`
	Password     string             `json:"password"`
	Address      struct {
		Province     string `json:"province"`
		District     string `json:"district"`
		Sub_district string `json:"sub_district"`
		Postal_code  string `json:"postal_code"`
		Street       string `json:"street"`
		Building     string `json:"building"`
		House_number string `json:"house_number"`
	} `json:"address"`
}
