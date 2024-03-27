package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// Seller model from seller collection
type Seller struct {
	ID          primitive.ObjectID   `json:"id" bson:"_id"`
	FirstName   string               `json:"first_name" bson:"first_name"`
	LastName    string               `json:"last_name" bson:"last_name"`
	Pets        []primitive.ObjectID `json:"pets"`
	BankAccount BankAccount          `json:"bank_account" bson:"bank_account"`
	License     string               `json:"license" bson:"license"`
	Status      string               `json:"status" bson:"status"`
}

// BankAccount model
type BankAccount struct {
	BankName          string `json:"bank_name" bson:"bank_name"`
	BankAccountName   string `json:"bank_account_name" bson:"bank_account_name"`
	BankAccountNumber string `json:"bank_account_number" bson:"bank_account_number"`
}
