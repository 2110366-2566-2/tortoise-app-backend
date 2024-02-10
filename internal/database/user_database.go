package database

import (
	"context"

	"github.com/2110366-2566-2/tortoise-app-backend/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
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
