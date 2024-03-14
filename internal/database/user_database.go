package database

import (
	"context"
	"fmt"

	"github.com/2110366-2566-2/tortoise-app-backend/internal/models"
	"github.com/2110366-2566-2/tortoise-app-backend/pkg/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// CheckUserExist is a function to check if the user already exists
func CheckUserExist(ctx context.Context, h *Handler, user models.User) (bool, error) {
	// Check if the user already exists
	filter := bson.M{"username": user.Username}
	err := h.db.Collection("users").FindOne(ctx, filter).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

// CheckEmailExist is a function to check if the email already exists
func CheckEmailExist(ctx context.Context, h *Handler, user models.User) (bool, error) {
	// Check if the email already exists
	filter := bson.M{"email": user.Email}
	err := h.db.Collection("users").FindOne(ctx, filter).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

// CreateUser is a function to create a new user
func CreateUser(ctx context.Context, h *Handler, user models.User) error {
	// Insert the user to the database
	_, err := h.db.Collection("users").InsertOne(ctx, user)
	if err != nil {
		return err
	}

	return nil
}

// GetUserByUserID returns a user by userID
func (h *Handler) GetUserByUserID(ctx context.Context, userID string) (*models.User, error) {
	var user models.User
	userObjID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return nil, fmt.Errorf("failed to convert petID to ObjectID: %v", err)
	}
	filter := bson.M{"_id": userObjID}
	err = h.db.Collection("users").FindOne(ctx, filter).Decode(&user)
	if err != nil {
		return nil, fmt.Errorf("failed to find user: %v", err)
	}
	return &user, nil
}

// UpdateOneUser updates a user
func (h *Handler) UpdateOneUser(ctx context.Context, userID string, data bson.M) (*mongo.SingleResult, error) {
	// Convert BSON M data to BSON B (bson.D)
	var updateDoc bson.D
	for k, v := range data {

		// Convert password to String and Hash the password
		if k == "password" {
			v = utils.HashPassword(v.(string))
			updateDoc = append(updateDoc, bson.E{Key: k, Value: v})
		} else if k != "old_password" {
			updateDoc = append(updateDoc, bson.E{Key: k, Value: v})
		}
	}
	// convert string to objID
	userObjID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return nil, fmt.Errorf("failed to convert userID to ObjectID: %v", err)
	}
	// return updated User
	res := h.db.Collection("users").FindOneAndUpdate(ctx, bson.M{"_id": userObjID}, bson.D{{Key: "$set", Value: updateDoc}}, options.FindOneAndUpdate().SetReturnDocument(options.After))
	if res.Err() != nil {
		return nil, fmt.Errorf("failed to update user: %v", res.Err())
	}

	return res, nil
}

// DeleteOneUser deletes a user
func (h *Handler) DeleteOneUser(ctx context.Context, userID string) (*mongo.DeleteResult, error) {

	// get UserId
	user, err := h.GetUserByUserID(ctx, userID)

	if err != nil {
		return nil, fmt.Errorf("failed to get user: %v", err)
	}

	// Delete user from users collection
	res, err := h.db.Collection("users").DeleteOne(ctx, bson.M{"_id": user.ID})
	if err != nil {
		return nil, fmt.Errorf("failed to delete user: %v", err)
	}

	// Delete pet from user's pets
	if user.Role == 1 {
		// delete user from seller collection
		h.db.Collection("sellers").DeleteOne(ctx, bson.M{"_id": user.ID})
		// delete user's pet from pet collection
		h.db.Collection("pets").DeleteMany(ctx, bson.M{"seller_id": user.ID})
	} else {
		//delete user from buyers
		h.db.Collection("buyers").DeleteOne(ctx, bson.M{"_id": user.ID})
	}
	return res, nil
}

func (h *Handler) GetUserByMail(ctx context.Context, data bson.M) (*models.User, error) {
	var user models.User
	filter := data
	err := h.db.Collection("users").FindOne(ctx, filter).Decode(&user)
	if err != nil {
		return nil, fmt.Errorf("failed to find user: %v", err)
	}
	return &user, nil
}
