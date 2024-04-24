package database

import (
	"context"

	"github.com/2110366-2566-2/tortoise-app-backend/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// MasterData methods

// GetAllMasterData returns all master data
func (h *Handler) GetAllMasterData(ctx context.Context) (*[]models.MasterData, error) {
	var master_opts = []bson.M{
		{"$group": bson.M{
			"_id": "$category",
			"species_count": bson.M{
				"$sum": bson.M{
					"$cond": bson.M{
						"if":   bson.M{"$isArray": "$species"},
						"then": bson.M{"$size": "$species"},
						"else": 0,
					},
				},
			},
			"species": bson.M{"$first": "$species"},
		}},
		{"$project": bson.M{
			"_id":           0,
			"category":      "$_id",
			"species_count": 1,
			"species":       1,
		}},
	}
	cursor, err := h.db.Collection("master_data").Aggregate(ctx, master_opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	var masterData []models.MasterData
	if err = cursor.All(ctx, &masterData); err != nil {
		return nil, err
	}
	return &masterData, nil
}

// GetMasterDataByCategory returns master data by Category
func (h *Handler) GetMasterDataByCategory(ctx context.Context, category string) (*models.MasterData, error) {
	var masterData models.MasterData
	filter := bson.M{"category": bson.M{"$regex": "^" + category + "$", "$options": "i"}}
	err := h.db.Collection("master_data").FindOne(ctx, filter).Decode(&masterData)
	if err != nil {
		return nil, err
	}
	masterData.SpeciesCount = len(masterData.Species)
	return &masterData, nil
}

// GetCategories returns list of categories
func (h *Handler) GetCategories(ctx context.Context) (*models.MasterDataCategory, error) {
	// project only category
	var opts = bson.A{
		bson.D{
			{Key: "$group", Value: bson.D{
				{Key: "_id", Value: primitive.Null{}},
				{Key: "categories", Value: bson.D{{Key: "$addToSet", Value: "$category"}}},
			}},
		},
		bson.D{
			{Key: "$project", Value: bson.D{
				{Key: "_id", Value: 0},
				{Key: "categories", Value: 1},
			}},
		},
	}

	var categories []models.MasterDataCategory
	cursor, err := h.db.Collection("master_data").Aggregate(ctx, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	if err = cursor.All(ctx, &categories); err != nil {
		return nil, err
	}
	return &categories[0], nil
}
