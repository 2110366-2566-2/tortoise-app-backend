package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// user model from user collection
type User struct {
	// ID represents the unique identifier of a user.
	ID          primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Username    string             `json:"username" binding:"required"` // unique
	Email       string             `json:"email" binding:"required"`    // unique
	Password    string             `json:"password" binding:"required"`
	FirstName   string             `json:"first_name" bson:"first_name"`
	LastName    string             `json:"last_name" bson:"last_name"`
	Gender      string             `json:"gender"`
	PhoneNumber string             `json:"phoneNumber" bson:"phoneNumber"`
	Image       string             `json:"image"`
	Role        int32              `json:"role" binding:"required"`
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
