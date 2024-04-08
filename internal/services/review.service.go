package services

import (
	"log"
	"net/http"
	"time"

	"github.com/2110366-2566-2/tortoise-app-backend/internal/database"
	"github.com/2110366-2566-2/tortoise-app-backend/internal/models"
	"github.com/2110366-2566-2/tortoise-app-backend/pkg/utils"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ReviewHandler struct {
	handler *database.Handler
}

func NewReviewHandler(handler *database.Handler) *ReviewHandler {
	return &ReviewHandler{handler: handler}
}

func (h *ReviewHandler) CreateReview(c *gin.Context) {
	var review models.Review
	if err := c.BindJSON(&review); err != nil {
		log.Println("Error: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "failed to bind Review"})
		return
	}

	review.Description = utils.SanitizeString(review.Description)

	// get user id from token
	userID, exist := c.Get("userID")
	if !exist {
		log.Println("Error: Unable to get userID")
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": "fail to authorize user"})
		return
	}

	review.Reviewer_id = userID.(primitive.ObjectID)
	review.Time = time.Now()

	_, err := h.handler.CreateReview(c, &review)
	if err != nil {
		log.Println("Error: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"success": true, "data": &review})
}

func (h *ReviewHandler) AddComment(c *gin.Context) {
	// get user id from token
	userID, exist := c.Get("userID")
	if !exist {
		log.Println("Error: Unable to get userID")
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": "fail to authorize user"})
		return
	}

	var comment models.Comments
	if err := c.BindJSON(&comment); err != nil {
		log.Println("Error: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "failed to bind comment"})
		return
	}
	comment.User_id = userID.(primitive.ObjectID)
	comment.Time = time.Now()

	res, err := h.handler.CreateComment(c, c.Param("reviewID"), comment)
	if err != nil {
		log.Println("Error: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": "failed to add comment"})
		return
	}
	// log result
	c.JSON(http.StatusOK, gin.H{"success": true, "data": res})
}

func (h *ReviewHandler) GetReviewBySeller(c *gin.Context) {
	reviews, err := h.handler.GetReviewByUserID(c, utils.SanitizeString(c.Param("sellerID")))
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

func (h *ReviewHandler) GetReviewByID(c *gin.Context) {
	review, err := h.handler.GetReviewByReviewID(c, utils.SanitizeString(c.Param("reviewID")))
	if err != nil {
		log.Println("Error: ", err)
		errorMsg := "failed to get review by id"
		if err.Error() == "review not found" {
			errorMsg = "review not found"
		}
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": errorMsg})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "data": review})
}

func (h *ReviewHandler) DeleteReview(c *gin.Context) {
	userID, exist := c.Get("userID")
	role, _ := c.Get("role")
	// check type of userID and role
	// fmt.Println("userID: ", userID)
	// fmt.Println("role: ", role)

	if !exist {
		log.Println("Error: Unable to get userID")
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": "fail to authorize user"})
		return
	}
	// c.JSON(http.StatusOK, gin.H{"success": true, "data": userID})
	res, err := h.handler.DeleteReview(c, c.Param("reviewID"), role.(string), userID.(primitive.ObjectID))
	if err != nil {
		log.Println("Error: ", err)
		if err.Error() == "unauthorized" {
			c.JSON(http.StatusForbidden, gin.H{"success": false, "error": "user is not authorized to delete review"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": "failed to delete review"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "data": res})
}
