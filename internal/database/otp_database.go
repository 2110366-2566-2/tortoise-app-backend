package database

import (
	"context"
	"time"

	"github.com/2110366-2566-2/tortoise-app-backend/internal/models"
	"go.mongodb.org/mongo-driver/bson"
)

func (h *Handler) GetOTPbyEmail(ctx context.Context, email string) (*models.OTP, error) {
	var otp models.OTP
	filter := bson.M{"email": email}
	err := h.db.Collection("otps").FindOne(ctx, filter).Decode(&otp)
	if err != nil {
		return nil, err
	}
	return &otp, nil
}

func (h *Handler) CreateOTP(ctx context.Context, otp, email string) error {
	// set TTL for OTP
	otpData := models.OTP{
		OTP:       otp,
		Email:     email,
		CreatedAt: time.Now(),
	}

	// check if the email already exists
	count, err := h.db.Collection("otps").CountDocuments(ctx, bson.M{"email": email})
	if err != nil {
		return err
	}

	if count > 0 {
		// delete the existing OTP
		err = h.DeleteOTP(ctx, email)
		if err != nil {
			return err
		}
	}

	_, err = h.db.Collection("otps").InsertOne(ctx, otpData)
	if err != nil {
		return err
	}
	return nil
}

func (h *Handler) DeleteOTP(ctx context.Context, email string) error {
	filter := bson.M{"email": email}
	_, err := h.db.Collection("otps").DeleteOne(ctx, filter)
	if err != nil {
		return err
	}
	return nil
}
