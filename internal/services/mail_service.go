package services

import (
	"log"
	"net/http"
	"net/smtp"

	"github.com/gin-gonic/gin"

	"go.mongodb.org/mongo-driver/bson"
)

func (h *UserHandler) Recovery_username(c *gin.Context) {
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
