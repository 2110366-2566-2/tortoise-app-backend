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
	"github.com/golang-jwt/jwt/v5"

	"time"

	"go.mongodb.org/mongo-driver/bson"
)

func (h *UserHandler) RecoveryUsername(c *gin.Context) {
	var data bson.M
	c.BindJSON(&data)

	// Sanitize data
	utils.BsonSanitize(&data)

	user, err := h.dbHandler.GetUserByMail(c, data)

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
	to := user.Email
	pass := "ruik nfvk adgj ncyu"
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	p1 := "You have requested to <em class=\"special\">recover the username</em> of your PetPal account."
	p2 := "Please find your <em>username</em> below:"
	p3 := "Username"
	body := "<html>" + utils.GenerateHTMLTemplate(user.Username, p1, p2, p3) + "</html>"
	text := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: Recovery Your Petpal Password\n"
	msg := []byte(text + mime + body)

	err = smtp.SendMail("smtp.gmail.com:587", smtp.PlainAuth("", from, pass, "smtp.gmail.com"), from, []string{to}, []byte(msg))

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"success": true, "data": "send successfully"})
}

func (h *UserHandler) SentOTP(c *gin.Context) {
	var data bson.M
	c.BindJSON(&data)

	// Sanitize data
	utils.BsonSanitize(&data)

	// Check is email exist
	user, err := h.dbHandler.GetUserByMail(c, data)
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
	pass := "ruik nfvk adgj ncyu"
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	p1 := "You have requested to <em class=\"special\">recover</em> your PetPal account."
	p2 := "Please find your <em>One Time Password (OTP)</em> below:"
	p3 := "OTP"
	body := "<html>" + utils.GenerateHTMLTemplate(otp, p1, p2, p3) + "</html>"
	text := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: Recovery Your Petpal Password\n"
	msg := []byte(text + mime + body)

	err = smtp.SendMail("smtp.gmail.com:587", smtp.PlainAuth("", from, pass, "smtp.gmail.com"), from, []string{to}, []byte(msg))

	if err != nil {
		log.Println("Error: ", err)
		c.JSON(500, gin.H{"error": "failed to send OTP"})
		return
	}

	log.Println("OTP: ", otp)

	// Remove old OTP in database
	// err = h.dbHandler.DeleteOTP(c, to)

	// Add OTP to database
	err = h.dbHandler.CreateOTP(c, hashOtp, to)
	if err != nil {
		log.Println("Error: ", err)
		c.JSON(500, gin.H{"success": false, "error": "failed to create OTP"})
		return
	}

	c.JSON(200, gin.H{"success": true, "data": "send OTP successfully"})
}

func (h *UserHandler) ValidateOTP(c *gin.Context) {
	var res models.OTPResponse
	err := c.BindJSON(&res)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "failed to bind data"})
		return
	}

	res.Email = utils.SanitizeString(res.Email)
	res.OTP = utils.SanitizeString(res.OTP)

	// Check is email exist
	user, err := h.dbHandler.GetUserByMail(c, bson.M{"email": res.Email})
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
	otp, err := h.dbHandler.GetOTPbyEmail(c, res.Email)
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

	// Add token
	var role string
	if user.Role == 1 {
		role = "seller"
	} else if user.Role == 2 {
		role = "buyer"
	}

	// Create a new token object, specifying signing method and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID":   user.ID,
		"username": user.Username,
		"role":     role,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte("secret"))
	if err != nil {
		c.JSON(500, gin.H{"error": "failed to generate token", "success": false})
		return
	}

	// Remove OTP
	if err = h.dbHandler.DeleteOTP(c, res.Email); err != nil {
		log.Println("Error: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": "failed to delete OTP"})
		return
	}

	c.JSON(200, gin.H{"success": true, "data": user, "token": tokenString})

}

func (h *UserHandler) CheckMail(c *gin.Context) {
	var data bson.M
	c.BindJSON(&data)

	utils.BsonSanitize(&data)

	user, err := h.dbHandler.GetUserByMail(c, data)

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
