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

// GetPetByPetID returns a pet by petID
func (h *Handler) GetPetByPetID(ctx context.Context, petID string) (*models.Pet, error) {
	ObjectID, err := primitive.ObjectIDFromHex(petID)
	if err != nil {
		return nil, fmt.Errorf("failed to convert petID to ObjectID: %v", err)
	}
	var pet models.Pet
	filter := bson.M{"_id": ObjectID}
	err = h.db.Collection("pets").FindOne(ctx, filter).Decode(&pet)
	if err != nil {
		return nil, fmt.Errorf("failed to find pet: %v", err)
	}
	return &pet, nil
}

// GetPetBySeller returns pets that belong to a seller
func (h *Handler) GetPetBySeller(ctx context.Context, userID string) (*[]models.Pet, error) {
	// Convert userID to ObjectID
	sellerID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return nil, fmt.Errorf("failed to convert userID to ObjectID: %v", err)
	}

	// Find pets by seller_id
	filter := bson.M{"seller_id": sellerID}
	cursor, err := h.db.Collection("pets").Find(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("failed to find pets: %v", err)
	}
	defer cursor.Close(ctx)

	// Fetch pets
	var pets []models.Pet
	for cursor.Next(ctx) {
		var pet models.Pet
		if err := cursor.Decode(&pet); err != nil {
			return nil, fmt.Errorf("failed to decode document: %v", err)
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
		return nil, fmt.Errorf("failed to convert userID to ObjectID: %v", err)
	}

	// Check if seller exists
	opts := options.FindOne().SetProjection(bson.M{"pets": 0})
	filter := bson.M{"_id": sellerID}
	seller := h.db.Collection("sellers").FindOne(ctx, filter, opts)
	if err := seller.Err(); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("seller not found")
		}
		return nil, fmt.Errorf("failed to find seller: %v", err)
	}

	// Insert pet to pets collection
	res, err := h.db.Collection("pets").InsertOne(ctx, pet)
	if err != nil {
		return nil, fmt.Errorf("failed to insert pet: %v", err)
	}

	// Update user's pets
	filter = bson.M{"_id": sellerID}
	update := bson.M{"$push": bson.M{"pets": res.InsertedID}}
	_, err = h.db.Collection("users").UpdateOne(ctx, filter, update)
	if err != nil {
		// rollback
		_, err2 := h.db.Collection("pets").DeleteOne(ctx, bson.M{"_id": res.InsertedID})
		if err2 != nil {
			return nil, fmt.Errorf("failed to rollback: %v", err2)
		}

		return nil, fmt.Errorf("failed to update user's pets: %v", err)
	}

	return res, nil

}

// UpdateOnePet updates a pet
func (h *Handler) UpdateOnePet(ctx context.Context, petID string, data bson.M) (*mongo.UpdateResult, error) {
	// Convert petID to ObjectID
	objectID, err := primitive.ObjectIDFromHex(petID)
	if err != nil {
		return nil, fmt.Errorf("failed to convert petID to ObjectID: %v", err)
	}

	// Convert BSON M data to BSON B (bson.D)
	var updateDoc bson.D
	for key, value := range data {
		updateDoc = append(updateDoc, bson.E{Key: key, Value: value})
	}

	// Update pet
	res, err := h.db.Collection("pets").UpdateOne(ctx, bson.M{"_id": objectID}, bson.D{{Key: "$set", Value: updateDoc}})
	if err != nil {
		return nil, fmt.Errorf("failed to update pet: %v", err)
	}

	return res, nil
}

// DeleteOnePet deletes a pet
func (h *Handler) DeleteOnePet(ctx context.Context, petID string) (*mongo.DeleteResult, error) {
	// Convert petID to ObjectID
	objectID, err := primitive.ObjectIDFromHex(petID)
	if err != nil {
		return nil, fmt.Errorf("failed to convert petID to objectID: %v", err)
	}
	// get sellerId
	pet, err := h.GetPetByPetID(ctx, petID)
	if err != nil {
		return nil, fmt.Errorf("failed to get pet: %v", err)
	}
	sellerID := pet.Seller_id

	// Delete pet from pets collection
	res, err := h.db.Collection("pets").DeleteOne(ctx, bson.M{"_id": objectID})
	if err != nil {
		return nil, fmt.Errorf("failed to delete pet: %v", err)
	}

	// Delete pet from user's pets
	filter := bson.M{"_id": sellerID}
	update := bson.M{"$pull": bson.M{"pets": objectID}}
	_, err = h.db.Collection("sellers").UpdateOne(ctx, filter, update)
	if err != nil {
		// rollback
		_, err2 := h.db.Collection("pets").InsertOne(ctx, pet)
		if err2 != nil {
			return nil, fmt.Errorf("failed to rollback: %v", err2)
		}

		return nil, fmt.Errorf("failed to delete pet from user's pets: %v", err)
	}

	return res, nil
}
