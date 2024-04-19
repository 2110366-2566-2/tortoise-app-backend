package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// user model from user collection
type User struct {
	// ID represents the unique identifier of a user.
	ID          primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Username    string             `json:"username" binding:"required" validate:"required,min=4,max=20"` // unique
	Email       string             `json:"email" binding:"required" validate:"required,email"`           // unique
	Password    string             `json:"password" binding:"required" validate:"required,min=6"`
	FirstName   string             `json:"first_name" bson:"first_name"`
	LastName    string             `json:"last_name" bson:"last_name"`
	Gender      string             `json:"gender"`
	PhoneNumber string             `json:"phone_number" bson:"phone_number"`
	Image       string             `json:"image"`
	Role        int32              `json:"role" binding:"required" validate:"required,eq=1|eq=2"`
	License     string             `json:"license,omitempty" bson:"license,omitempty"`
	Address     struct {
		Province    string `json:"province"`
		District    string `json:"district"`
		SubDistrict string `json:"subdistrict"`
		PostalCode  string `json:"postalCode"`
		Street      string `json:"street"`
		Building    string `json:"building"`
		HouseNumber string `json:"houseNumber"`
	} `json:"address"`
}

// Password model
type Password struct {
	Password string `json:"password" binding:"required"`
}
