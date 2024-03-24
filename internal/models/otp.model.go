package models

import "time"

// OTP model
type OTP struct {
	ID        string    `json:"id,omitempty" bson:"_id,omitempty"`
	Email     string    `json:"email" bson:"email" binding:"required"`
	OTP       string    `json:"otp" bson:"otp" binding:"required"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
}

type OTPResponse struct {
	Email string `json:"email" binding:"required"`
	OTP   string `json:"otp" binding:"required"`
}

type ResetPassword struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
