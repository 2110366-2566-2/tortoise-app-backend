package services

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"net/smtp"

	"github.com/2110366-2566-2/tortoise-app-backend/internal/models"
	"github.com/2110366-2566-2/tortoise-app-backend/pkg/utils"
	"github.com/gin-gonic/gin"

	"time"

	"go.mongodb.org/mongo-driver/bson"
)

func (h *UserHandler) RecoveryUsername(c *gin.Context) {
	var data bson.M
	c.BindJSON(&data)

	user, err := h.handler.GetUserByMail(c, data)

	if err != nil {
		log.Println("Error: ", err)
		errorMsg := "failed to get user by user id"
		if err.Error() == "mongo: no documents in result" {
			errorMsg = "user not found"
		}
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": errorMsg})
		return
	}

	from := "petpal.tortoise@gmail.com"
	pass := "secl pvjq qpsv jynu"
	to := user.Email
	body := "Your Username is " + user.Username

	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: Rcovery Your Petpal Username\n\n" +
		body

	err = smtp.SendMail("smtp.gmail.com:587", smtp.PlainAuth("", from, pass, "smtp.gmail.com"), from, []string{to}, []byte(msg))

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"success": true, "data": "send successfully"})
}

// SentOTP godoc
// @Method POST
// @Summary Sent OTP
// @Description Sent OTP to user's email
// @Endpoint /api/v1/user/sent-otp
func (h *UserHandler) SentOTP(c *gin.Context) {
	var data bson.M
	c.BindJSON(&data)

	// Check is email exist
	user, err := h.handler.GetUserByMail(c, data)
	if err != nil {
		log.Println("Error: ", err)
		errorMsg := "failed to get user by user id"
		if err.Error() == "mongo: no documents in result" {
			errorMsg = "user not found"
		}
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": errorMsg})
		return
	}

	to := user.Email

	rand.Seed(time.Now().UnixNano())
	otp := fmt.Sprintf("%06d", rand.Intn(999999))

	hashOtp := utils.HashPassword(otp)

	from := "petpal.tortoise@gmail.com"
	pass := "secl pvjq qpsv jynu"
	body := "The OTP is " + otp
	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: Rcovery Your Petpal Password\n\n" +
		body

	err = smtp.SendMail("smtp.gmail.com:587", smtp.PlainAuth("", from, pass, "smtp.gmail.com"), from, []string{to}, []byte(msg))

	if err != nil {
		log.Println("Error: ", err)
		c.JSON(500, gin.H{"error": "failed to send OTP"})
		return
	}

	log.Println("OTP: ", otp)

	// Remove old OTP in database
	// err = h.handler.DeleteOTP(c, to)

	// Add OTP to database
	err = h.handler.CreateOTP(c, hashOtp, to)
	if err != nil {
		log.Println("Error: ", err)
		c.JSON(500, gin.H{"success": false, "error": "failed to create OTP"})
		return
	}

	c.JSON(200, gin.H{"success": true, "data": "send OTP successfully"})
}

// ValidateOTP godoc
// @Method POST
// @Summary Validate OTP
// @Description Validate OTP
// @Endpoint /api/v1/user/checkotp
func (h *UserHandler) ValidateOTP(c *gin.Context) {
	var res models.OTPResponse
	err := c.BindJSON(&res)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "failed to bind data"})
		return
	}

	// Check is email exist
	user, err := h.handler.GetUserByMail(c, bson.M{"email": res.Email})
	if err != nil {
		log.Println("Error: ", err)
		errorMsg := "failed to get user by user id"
		if err.Error() == "mongo: no documents in result" {
			errorMsg = "user not found"
		}
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": errorMsg})
		return
	}

	// Check if the OTP is valid
	otp, err := h.handler.GetOTPbyEmail(c, res.Email)
	if err != nil {
		log.Println("Error: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "OTP not found"})
		return
	}

	// Check if the OTP is valid
	if !utils.ComparePassword(otp.OTP, res.OTP) {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "OTP is incorrect"})
		return
	}

	c.JSON(200, gin.H{"success": true, "data": user})

}

func (h *UserHandler) CheckMail(c *gin.Context) {
	var data bson.M
	c.BindJSON(&data)

	user, err := h.handler.GetUserByMail(c, data)

	if err != nil {
		log.Println("Error: ", err)
		errorMsg := "failed to get user by user id"
		if err.Error() == "mongo: no documents in result" {
			errorMsg = "user not found"
		}
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": errorMsg})
		return
	}

	c.JSON(200, gin.H{"success": true, "data": user})
}
