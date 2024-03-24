package apiV1

import (
	"github.com/2110366-2566-2/tortoise-app-backend/internal/database"
	"github.com/2110366-2566-2/tortoise-app-backend/internal/services"
	"github.com/gin-gonic/gin"
)

func ReviewServices(r *gin.RouterGroup, h *database.Handler) {

	// Create a new review handler
	reviewHandler := services.NewReviewHandler(h)

	// Set up routes
	r.POST("/create", reviewHandler.CreateReview)
	r.PUT("/comment/:reviewID", reviewHandler.AddComment)
	r.GET("/seller/:sellerID", reviewHandler.GetReviewBySeller)
	r.GET("/:reviewID", reviewHandler.GetReviewByID)
	r.DELETE("/:reviewID", reviewHandler.DeleteReview)
}
