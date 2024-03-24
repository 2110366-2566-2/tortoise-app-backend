package apiV1

import (
	"net/http"

	"github.com/2110366-2566-2/tortoise-app-backend/internal/database"
	"github.com/2110366-2566-2/tortoise-app-backend/internal/services"
	"github.com/gin-gonic/gin"
)

func AdminServices(r *gin.RouterGroup, h *database.Handler) {
	// Create a new admin handler
	admin_report_handler := services.NewReportHandler(h)

	// Set up routes
	r.GET("/report", admin_report_handler.GetReport)
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "This is admin service",
		})
	})

	// Approve seller
	r.POST("/approve-seller/:sellerID", func(c *gin.Context) {
		services.ApproveSeller(c, h)
	})

}
