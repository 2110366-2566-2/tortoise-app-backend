package database

import (
	"context"

	"github.com/2110366-2566-2/tortoise-app-backend/internal/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// CheckAdminExist is a function to check if the Admin already exists
func CheckAdminExist(ctx context.Context, h *Handler, admin models.Admin) (bool, error) {
	// Check if the Admin already exists
	filter := bson.M{"username": admin.Username}
	err := h.db.Collection("admins").FindOne(ctx, filter).Decode(&admin)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

// CheckEmailExist is a function to check if the email already exists
func CheckAdminEmailExist(ctx context.Context, h *Handler, admin models.Admin) (bool, error) {
	// Check if the email already exists
	filter := bson.M{"email": admin.Email}
	err := h.db.Collection("admins").FindOne(ctx, filter).Decode(&admin)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

// CreateAdmin is a function to create a new Admin
func CreateAdmin(ctx context.Context, h *Handler, admin models.Admin) error {
	// Insert the Admin to the database
	_, err := h.db.Collection("admins").InsertOne(ctx, admin)
	if err != nil {
		return err
	}

	return nil
}

// GetAdminByUserID is a function to get an Admin by userID
func GetAdminByUserID(ctx context.Context, h *Handler, userID primitive.ObjectID) (*models.Admin, error) {
	var admin models.Admin

	filter := bson.M{"_id": userID}

	err := h.db.Collection("admins").FindOne(ctx, filter).Decode(&admin)
	if err != nil {
		return nil, err
	}

	return &admin, nil
}

// GetAdminByUsername is a function to get an Admin by username
func GetAdminByUsername(ctx context.Context, h *Handler, username string) (*models.Admin, error) {
	var admin models.Admin

	filter := bson.M{"username": username}

	err := h.db.Collection("admins").FindOne(ctx, filter).Decode(&admin)
	if err != nil {
		return nil, err
	}

	return &admin, nil
}

// GetAdminByEmail is a function to get an Admin by email
func GetAdminByEmail(ctx context.Context, h *Handler, email string) (*models.Admin, error) {
	var admin models.Admin

	filter := bson.M{"email": email}

	err := h.db.Collection("admins").FindOne(ctx, filter).Decode(&admin)
	if err != nil {
		return nil, err
	}

	return &admin, nil
}
