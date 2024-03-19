package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// admin model from user collection
type Admin struct {
	ID          primitive.ObjectID `json:"id" bson:"_id"`
	Username    string `json:"username" binding:"required"` // unique
	Email       string `json:"email" binding:"required"`    // unique
	Password    string `json:"password" binding:"required"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Gender      string `json:"gender"`
	PhoneNumber string `json:"phoneNumber"`
	Image       string `json:"image"`
	// Role        int32  `json:"role" binding:"required"`
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
