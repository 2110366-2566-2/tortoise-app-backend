package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// admin model from user collection
type Admin struct {
	ID          primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Username    string             `json:"username" binding:"required"` // unique
	Email       string             `json:"email" binding:"required"`    // unique
	Password    string             `json:"password" binding:"required"`
	FirstName   string             `json:"first_name" bson:"first_name" binding:"required"`
	LastName    string             `json:"last_name" bson:"last_name" binding:"required"`
	Gender      string             `json:"gender"`
	PhoneNumber string             `json:"phone_number" bson:"phone_number" binding:"required"`
	Image       string             `json:"image"`
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
