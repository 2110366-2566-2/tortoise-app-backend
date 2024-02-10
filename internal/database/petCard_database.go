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
func (h *Handler) GetAllPetCards(ctx context.Context) (*[]models.PetCard, error) {
	// Define the pipeline

	pipeline := mongo.Pipeline{
		bson.D{{Key: "$match", Value: bson.D{{Key: "is_sold", Value: false}}}},
		bson.D{
			{Key: "$lookup",
				Value: bson.D{
					{Key: "from", Value: "sellers"},
					{Key: "localField", Value: "seller_id"},
					{Key: "foreignField", Value: "_id"},
					{Key: "as", Value: "seller"},
				},
			},
		},
		bson.D{{Key: "$unwind", Value: "$seller"}},
		bson.D{{Key: "$project",
			Value: bson.D{
				{Key: "_id", Value: 1},
				{Key: "name", Value: 1},
				{Key: "type", Value: 1},
				{Key: "species", Value: 1},
				{Key: "price", Value: 1},
				{Key: "media", Value: 1},
				{Key: "seller_id", Value: 1},
				{Key: "seller_name", Value: "$seller.first_name"},
				{Key: "seller_surname", Value: "$seller.last_name"},
			},
		},
		},
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

	return &petCards, nil
}
