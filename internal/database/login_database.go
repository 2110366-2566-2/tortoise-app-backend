package database

import (
	"fmt"
	"context"
	"github.com/2110366-2566-2/tortoise-app-backend/internal/models"
	"github.com/2110366-2566-2/tortoise-app-backend/pkg/utils"
	"go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
)

func ValidateLogin(ctx context.Context, h *Handler, username string, password string) (*models.User, error) {
	var user models.User

	filter := bson.M{"username": username}
	err := h.db.Collection("users").FindOne(ctx, filter).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("user not found")
		}
		return nil, err
	}

	if !utils.ComparePassword(user.Password, password) {
		return nil, fmt.Errorf("invalid password")
	}

	return &user, nil
}