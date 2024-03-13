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
	_, err := h.db.Collection("otps").InsertOne(ctx, otpData)
	if err != nil {
		return err
	}
	return nil
}
