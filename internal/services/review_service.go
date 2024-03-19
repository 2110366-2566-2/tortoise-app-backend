package services

import (
	"log"
	"net/http"

	"github.com/2110366-2566-2/tortoise-app-backend/internal/database"
	"github.com/2110366-2566-2/tortoise-app-backend/internal/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

type ReviewHandler struct {
	handler *database.Handler
}

func NewReviewHandler(handler *database.Handler) *ReviewHandler {
	return &ReviewHandler{handler: handler}
}

// CreateReview godoc
// @Method POST
// @Summary Create review
// @Description Create review with user id
// @Endpoint /api/v1/review/create
func (h *ReviewHandler) CreateReview(c *gin.Context) {
	var review models.Review
	if err := c.BindJSON(&review); err != nil {
		log.Println("Error: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "failed to bind Review"})
		return
	}
	_, err := h.handler.CreateReview(c, &review)
	if err != nil {
		log.Println("Error: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"success": true, "data": &review})
}

func (h *ReviewHandler) AddComment(c *gin.Context) {
	var data bson.M
	c.BindJSON(&data)

	res, err := h.handler.CreateComment(c, c.Param("reviewID"), data)
	if err != nil {
		log.Println("Error: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": "failed to add comment"})
		return
	}
	// log result
	c.JSON(http.StatusOK, gin.H{"success": true, "data": res})
}

func (h *ReviewHandler) GetReviewBySeller(c *gin.Context) {
	reviews, err := h.handler.GetReviewByUserID(c, c.Param("sellerID"))
	if err != nil {
		log.Println("Error: ", err)
		errorMsg := "failed to get reviews by seller"
		if err.Error() == "seller not found" {
			errorMsg = "seller not found"
		}
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": errorMsg})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "count": len(*reviews), "data": &reviews})
}
