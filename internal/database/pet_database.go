package database

import (
	"context"
	"fmt"

	"github.com/2110366-2566-2/tortoise-app-backend/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Pet methods

// GetAllPets returns all pets
func (h *Handler) GetAllPetCards(ctx context.Context) (*[]models.PetCard, error) {
	var pets []models.PetCard
	pipeline := mongo.Pipeline{
		{{Key: "$match", Value: bson.M{"is_sold": false}}},
		{{Key: "$lookup", Value: bson.M{
			"from":         "sellers",
			"localField":   "seller_id",
			"foreignField": "_id",
			"as":           "seller",
		}}},
		{{Key: "$unwind", Value: "$seller"}},
		{{Key: "$project", Value: bson.M{
			"_id":            1,
			"name":           1,
			"type":           1,
			"price":          1,
			"media":          1,
			"seller_id":      1,
			"seller_name":    "$seller.name",
			"seller_surname": "$seller.surname",
		}}},
	}
	cursor, err := h.db.Collection("pets").Aggregate(ctx, pipeline, options.Aggregate())
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var pet models.PetCard
		if err := cursor.Decode(&pet); err != nil {
			return nil, err
		}
		pets = append(pets, pet)
	}
	return &pets, nil
}

// GetPetByPetID returns a pet by petID
func (h *Handler) GetPetByPetID(ctx context.Context, petID string) (*models.Pet, error) {
	var pet models.Pet
	ObjectID, err := primitive.ObjectIDFromHex(petID)
	if err != nil {
		return nil, err
	}
	err = h.db.Collection("pets").FindOne(ctx, map[string]interface{}{"_id": ObjectID}).Decode(&pet)
	if err != nil {
		return nil, err
	}
	return &pet, nil
}

// GetPetBySeller returns pets that belong to a seller
func (h *Handler) GetPetBySeller(ctx context.Context, userID string) (*[]models.Pet, error) {
	var pets []models.Pet
	// Convert userID to ObjectID
	sellerID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return nil, err
	}
	// Find pets by seller_id
	cursor, err := h.db.Collection("pets").Find(ctx, map[string]interface{}{"seller_id": sellerID.Hex()})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	// Fetch pets
	for cursor.Next(ctx) {
		var pet models.Pet
		if err := cursor.Decode(&pet); err != nil {
			return nil, err
		}
		pets = append(pets, pet)
	}
	return &pets, nil
}

// CreateOnePet creates a new pet
func (h *Handler) CreateOnePet(ctx context.Context, userID string, pet *models.Pet) (*mongo.InsertOneResult, error) {
	// Update user's pets
	sellerID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return nil, err
	}
	// find if user exists
	count, err := h.db.Collection("users").CountDocuments(ctx, bson.M{"_id": sellerID, "role": 1})
	if err != nil {
		return nil, err
	}
	if count == 0 {
		return nil, fmt.Errorf("seller not found")
	}

	// Insert pet to pets collection
	res, err := h.db.Collection("pets").InsertOne(ctx, pet)
	if err != nil {
		return nil, err
	}

	// Update user's pets
	_, err = h.db.Collection("users").UpdateOne(ctx, map[string]interface{}{"_id": sellerID}, map[string]interface{}{"$push": map[string]interface{}{"pets": res.InsertedID}})
	if err != nil {
		return nil, err
	}

	return res, nil

}

// func (h *Handler) UpdateOnePet(ctx context.Context, petID string, edit map[string]string) (*mongo.UpdateResult, error) {
// 	ObjectID, err := primitive.ObjectIDFromHex(petID)
// 	if err != nil {
// 		return nil, err
// 	}
// 	res, err := h.db.Collection("pets").UpdateOne(ctx, map[string]interface{}{"_id": ObjectID}, edit)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return res, nil
// }
