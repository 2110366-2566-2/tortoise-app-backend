package database

import (
	"context"
	// "fmt"

	"github.com/2110366-2566-2/tortoise-app-backend/internal/models"
	// "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	// "go.mongodb.org/mongo-driver/mongo"
)

// User methods

// GetAllUsers returns all users
func (h *Handler) GetAllUsers(ctx context.Context) (*[]models.User, error) {
	var users []models.User
	cursor, err := h.db.Collection("users").Find(ctx, map[string]string{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var user models.User
		if err := cursor.Decode(&user); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return &users, nil
}

// GetUserByUserID returns a user by userID
func (h *Handler) GetUserByUserID(ctx context.Context, userID string) (*models.User, error) {
	var user models.User
	ObjectID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return nil, err
	}
	err = h.db.Collection("users").FindOne(ctx, map[string]interface{}{"_id": ObjectID}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}


