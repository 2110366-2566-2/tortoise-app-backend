package database

import (
	"context"
	"fmt"

	"github.com/2110366-2566-2/tortoise-app-backend/internal/models"
	"github.com/2110366-2566-2/tortoise-app-backend/pkg/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func ValidateLogin(ctx context.Context, h *Handler, username string, password string) (*models.User, error) {
	var user models.User

	filter := bson.M{"username": username}

	err := h.db.Collection("users").FindOne(ctx, filter).Decode(&user)
	if err == nil && utils.ComparePassword(user.Password, password) {
		return &user, nil
	}

	if err == mongo.ErrNoDocuments {
		return nil, fmt.Errorf("user not found")
	}

	return nil, fmt.Errorf("invalid username or password")
}

func AdminValidateLogin(ctx context.Context, h *Handler, username string, password string) (*models.Admin, error) {
	var admin models.Admin

	filter := bson.M{"username": username}

	err := h.db.Collection("admins").FindOne(ctx, filter).Decode(&admin)
	if err == nil && utils.ComparePassword(admin.Password, password) {
		return &admin, nil
	}

	if err == mongo.ErrNoDocuments {
		return nil, fmt.Errorf("admin not found")
	}

	return nil, fmt.Errorf("invalid username or password")
}
