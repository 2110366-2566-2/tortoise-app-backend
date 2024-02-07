package database

import (
	"context"
	"fmt"

	"github.com/2110366-2566-2/tortoise-app-backend/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// PetCard methods

// GetAllPetsCard returns all pets with some fields and seller's name and surname
func (h *Handler) GetAllPetCards(ctx context.Context) ([]models.PetCard, error) {
	// Define the pipeline
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

	// Execute aggregation
	cursor, err := h.db.Collection("pets").Aggregate(ctx, pipeline)
	if err != nil {
		return nil, fmt.Errorf("failed to aggregate: %v", err)
	}
	defer cursor.Close(ctx)

	// Decode results
	var petCards []models.PetCard
	for cursor.Next(ctx) {
		var petCard models.PetCard
		if err := cursor.Decode(&petCard); err != nil {
			return nil, fmt.Errorf("failed to decode document: %v", err)
		}
		petCards = append(petCards, petCard)
	}
	if err := cursor.Err(); err != nil {
		return nil, fmt.Errorf("cursor error: %v", err)
	}

	return petCards, nil
}
