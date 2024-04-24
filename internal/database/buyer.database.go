package database

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (h *Handler) UpdateBuyer(ctx context.Context, buyerID string, data bson.M) (*mongo.UpdateResult, error) {
	// Convert buyerID to ObjectID
	buyerObjID, err := primitive.ObjectIDFromHex(buyerID)
	if err != nil {
		return nil, fmt.Errorf("failed to convert buyerID to ObjectID: %v", err)
	}

	// Update buyer
	filter := bson.M{"_id": buyerObjID}
	update := bson.M{"$set": data}
	res, err := h.db.Collection("buyers").UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, fmt.Errorf("failed to update buyer: %v", err)
	}
	return res, nil
}
