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
// func (h *Handler) GetAllPetCards(ctx context.Context) (*[]models.PetCard, error) {
// 	// Define the pipeline

// 	pipeline := mongo.Pipeline{
// 		bson.D{{Key: "$match", Value: bson.D{{Key: "is_sold", Value: false}}}},
// 		bson.D{
// 			{Key: "$lookup",
// 				Value: bson.D{
// 					{Key: "from", Value: "sellers"},
// 					{Key: "localField", Value: "seller_id"},
// 					{Key: "foreignField", Value: "_id"},
// 					{Key: "as", Value: "seller"},
// 				},
// 			},
// 		},
// 		bson.D{{Key: "$unwind", Value: "$seller"}},
// 		bson.D{{Key: "$project",
// 			Value: bson.D{
// 				{Key: "_id", Value: 1},
// 				{Key: "name", Value: 1},
// 				{Key: "category", Value: 1},
// 				{Key: "species", Value: 1},
// 				{Key: "price", Value: 1},
// 				{Key: "media", Value: 1},
// 				{Key: "seller_id", Value: 1},
// 				{Key: "seller_name", Value: "$seller.first_name"},
// 				{Key: "seller_surname", Value: "$seller.last_name"},
// 			},
// 		},
// 		},
// 	}

// 	// Execute aggregation
// 	cursor, err := h.db.Collection("pets").Aggregate(ctx, pipeline)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to aggregate: %v", err)
// 	}
// 	defer cursor.Close(ctx)

// 	// Decode results
// 	var petCards []models.PetCard
// 	for cursor.Next(ctx) {
// 		var petCard models.PetCard
// 		if err := cursor.Decode(&petCard); err != nil {
// 			return nil, fmt.Errorf("failed to decode document: %v", err)
// 		}
// 		petCards = append(petCards, petCard)
// 	}
// 	if err := cursor.Err(); err != nil {
// 		return nil, fmt.Errorf("cursor error: %v", err)
// 	}

// 	return &petCards, nil
// }

// Get filtered petCard
func (h *Handler) GetFilteredPetCards(ctx context.Context, categories, species, sex, behaviors []string, minAge, maxAge, minWeight, maxWeight, minPrice, maxPrice int) (*[]models.PetCard, error) {
	// Define the pipeline
	pipeline := mongo.Pipeline{
		bson.D{{Key: "$match", Value: bson.D{{Key: "is_sold", Value: false}}}},
	}
	if len(categories) > 0 {
		pipeline = append(pipeline, bson.D{{Key: "$match", Value: bson.D{{Key: "category", Value: bson.D{{Key: "$in", Value: categories}}}}}})
	}
	if len(species) > 0 {
		pipeline = append(pipeline, bson.D{{Key: "$match", Value: bson.D{{Key: "species", Value: bson.D{{Key: "$in", Value: species}}}}}})
	}
	if len(sex) > 0 {
		pipeline = append(pipeline, bson.D{{Key: "$match", Value: bson.D{{Key: "sex", Value: bson.D{{Key: "$in", Value: sex}}}}}})
	}
	if len(behaviors) > 0 {
		pipeline = append(pipeline, bson.D{{Key: "$match", Value: bson.D{{Key: "behavior", Value: bson.D{{Key: "$in", Value: behaviors}}}}}})
	}
	if minAge != 0 || maxAge != 0 {
		ageFilter := bson.D{}
		if minAge != 0 {
			ageFilter = append(ageFilter, bson.E{Key: "$gte", Value: minAge})
		}
		if maxAge != 0 {
			ageFilter = append(ageFilter, bson.E{Key: "$lte", Value: maxAge})
		}
		pipeline = append(pipeline, bson.D{{Key: "$match", Value: bson.D{{Key: "age", Value: ageFilter}}}})
	}
	if minWeight != 0 || maxWeight != 0 {
		weightFilter := bson.D{}
		if minWeight != 0 {
			weightFilter = append(weightFilter, bson.E{Key: "$gte", Value: minWeight})
		}
		if maxWeight != 0 {
			weightFilter = append(weightFilter, bson.E{Key: "$lte", Value: maxWeight})
		}
		pipeline = append(pipeline, bson.D{{Key: "$match", Value: bson.D{{Key: "weight", Value: weightFilter}}}})
	}
	if minPrice != 0 || maxPrice != 0 {
		priceFilter := bson.D{}
		if minPrice != 0 {
			priceFilter = append(priceFilter, bson.E{Key: "$gte", Value: minPrice})
		}
		if maxPrice != 0 {
			priceFilter = append(priceFilter, bson.E{Key: "$lte", Value: maxPrice})
		}
		pipeline = append(pipeline, bson.D{{Key: "$match", Value: bson.D{{Key: "price", Value: priceFilter}}}})
	}

	pipeline = append(pipeline, bson.D{
		{Key: "$lookup",
			Value: bson.D{
				{Key: "from", Value: "sellers"},
				{Key: "localField", Value: "seller_id"},
				{Key: "foreignField", Value: "_id"},
				{Key: "as", Value: "seller"},
			},
		},
	})
	pipeline = append(pipeline, bson.D{{Key: "$unwind", Value: "$seller"}})
	pipeline = append(pipeline, bson.D{{Key: "$project",
		Value: bson.D{
			{Key: "_id", Value: 1},
			{Key: "name", Value: 1},
			{Key: "category", Value: 1},
			{Key: "species", Value: 1},
			{Key: "price", Value: 1},
			{Key: "media", Value: 1},
			{Key: "seller_id", Value: 1},
			{Key: "seller_name", Value: "$seller.first_name"},
			{Key: "seller_surname", Value: "$seller.last_name"},
		},
	},
	})

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
			return nil, fmt.Errorf("failed to decode document: %v", err) //return error
		}
		petCards = append(petCards, petCard)
	}
	if err := cursor.Err(); err != nil {
		return nil, fmt.Errorf("cursor error: %v", err)
	}

	return &petCards, nil
}
