package services

import (
	"log"
	"net/http"

	"github.com/2110366-2566-2/tortoise-app-backend/internal/database"
	"github.com/2110366-2566-2/tortoise-app-backend/internal/models"
	"github.com/gin-gonic/gin"
	// "go.mongodb.org/mongo-driver/bson"
)

type ReportHandler struct {
	handler *database.Handler
}

func NewReportHandler(handler *database.Handler) *ReportHandler {
	return &ReportHandler{handler: handler}
}

// func (h *ReviewHandler) CreateReview(c *gin.Context) {
// 	var review models.Review
// 	if err := c.BindJSON(&review); err != nil {
// 		log.Println("Error: ", err)
// 		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "failed to bind Review"})
// 		return
// 	}
// 	_, err := h.handler.CreateReview(c, &review)
// 	if err != nil {
// 		log.Println("Error: ", err)
// 		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
// 		return
// 	}

// 	c.JSON(http.StatusCreated, gin.H{"success": true, "data": &review})
// }

func (h *ReportHandler) CreatePartyReport(c *gin.Context) {
	var report models.PartyReport
	if err := c.BindJSON(&report); err != nil {
		log.Println("Error: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "failed to bind Report"})
		return
	}
	_, err := h.handler.CreatePartyReport(c, &report)
	if err != nil {
		log.Println("Error: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"success": true, "data": &report})
}

func (h *ReportHandler) CreateSystemReport(c *gin.Context) {
	var report models.SystemReport
	if err := c.BindJSON(&report); err != nil {
		log.Println("Error: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "failed to bind partyReport"})
		return
	}
	_, err := h.handler.CreateSystemReport(c, &report)
	if err != nil {
		log.Println("Error: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"success": true, "data": &report})
}
// func (h *ReviewHandler) AddComment(c *gin.Context) {
// 	var data bson.M
// 	c.BindJSON(&data)

// 	res, err := h.handler.CreateComment(c, c.Param("reviewID"), data)
// 	if err != nil {
// 		log.Println("Error: ", err)
// 		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": "failed to add comment"})
// 		return
// 	}
// 	// log result
// 	c.JSON(http.StatusOK, gin.H{"success": true, "data": res})
// }

// func (h *ReviewHandler) GetReviewBySeller(c *gin.Context) {
// 	reviews, err := h.handler.GetReviewByUserID(c, c.Param("sellerID"))
// 	if err != nil {
// 		log.Println("Error: ", err)
// 		errorMsg := "failed to get reviews by seller"
// 		if err.Error() == "seller not found" {
// 			errorMsg = "seller not found"
// 		}
// 		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": errorMsg})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"success": true, "count": len(*reviews), "data": &reviews})
// }
