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
	var pet models.Pet
	petObjID, err := primitive.ObjectIDFromHex(petID)
	if err != nil {
		return nil, fmt.Errorf("failed to convert petID to ObjectID: %v", err)
	}
	filter := bson.M{"_id": petObjID}
	err = h.db.Collection("pets").FindOne(ctx, filter).Decode(&pet)
	if err != nil {
		return nil, fmt.Errorf("failed to find pet: %v", err)
	}
	return &pet, nil
}

// GetPetBySeller returns pets that belong to a seller
func (h *Handler) GetPetBySeller(ctx context.Context, userID string) (*[]models.Pet, error) {
	// Convert string to ObjectID
	userObjID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return nil, fmt.Errorf("failed to convert userID to ObjectID: %v", err)
	}
	// Check if the seller exists
	sellerFilter := bson.M{"_id": userObjID}
	seller := h.db.Collection("sellers").FindOne(ctx, sellerFilter)
	if seller.Err() != nil {
		if seller.Err() == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("seller not found")
		}
		return nil, fmt.Errorf("failed to find seller: %v", seller.Err())
	}

	// Find pets for the given seller
	petFilter := bson.M{"seller_id": userObjID}
	cursor, err := h.db.Collection("pets").Find(ctx, petFilter)
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
func (h *Handler) CreateOnePet(ctx context.Context, userID primitive.ObjectID, pet *models.Pet) error {
	// Check if seller exists
	opts := options.FindOne().SetProjection(bson.M{"pets": 0})
	filter := bson.M{"_id": userID}
	seller := h.db.Collection("sellers").FindOne(ctx, filter, opts)
	if err := seller.Err(); err != nil {
		return fmt.Errorf("failed to find seller")
	}

	// Insert pet to pets collection
	_, err := h.db.Collection("pets").InsertOne(ctx, pet)
	if err != nil {
		if err.Error()[32:38] == "E11000" {
			return fmt.Errorf("pet already exists")
		}
		return fmt.Errorf("failed to insert pet")
	}

	// Update user's pets
	filter = bson.M{"_id": userID}
	update := bson.M{"$push": bson.M{"pets": pet.ID}}
	_, err = h.db.Collection("sellers").UpdateOne(ctx, filter, update)
	if err != nil {
		// rollback
		_, err2 := h.db.Collection("pets").DeleteOne(ctx, bson.M{"_id": pet.ID})
		if err2 != nil {
			return fmt.Errorf("failed to rollback: %v", err2)
		}

		return fmt.Errorf("failed to update user's pets: %v", err)
	}

	return nil

}

// UpdateOnePet updates a pet
func (h *Handler) UpdateOnePet(ctx context.Context, petID string, data bson.M) (*mongo.SingleResult, error) {
	// Convert BSON M data to BSON B (bson.D)
	var updateDoc bson.D
	for k, v := range data {
		updateDoc = append(updateDoc, bson.E{Key: k, Value: v})
	}
	// convert string to objID
	petObjID, err := primitive.ObjectIDFromHex(petID)
	if err != nil {
		return nil, fmt.Errorf("failed to convert petID to ObjectID: %v", err)
	}
	// return updated pet
	res := h.db.Collection("pets").FindOneAndUpdate(ctx, bson.M{"_id": petObjID}, bson.D{{Key: "$set", Value: updateDoc}}, options.FindOneAndUpdate().SetReturnDocument(options.After))
	if res.Err() != nil {
		return nil, fmt.Errorf("failed to update pet: %v", res.Err())
	}

	return res, nil
}

// DeleteOnePet deletes a pet
func (h *Handler) DeleteOnePet(ctx context.Context, petID string) (*mongo.DeleteResult, error) {
	// get sellerId
	pet, err := h.GetPetByPetID(ctx, petID)
	if err != nil {
		return nil, fmt.Errorf("failed to get pet: %v", err)
	}

	// Delete pet from pets collection
	res, err := h.db.Collection("pets").DeleteOne(ctx, bson.M{"_id": pet.ID})
	if err != nil {
		return nil, fmt.Errorf("failed to delete pet: %v", err)
	}

	// Delete pet from user's pets
	filter := bson.M{"_id": pet.Seller_id}
	update := bson.M{"$pull": bson.M{"pets": pet.ID}}
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
