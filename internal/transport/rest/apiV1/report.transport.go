package apiV1

import (
	"github.com/2110366-2566-2/tortoise-app-backend/internal/database"
	"github.com/2110366-2566-2/tortoise-app-backend/internal/services"
	"github.com/gin-gonic/gin"
)

func ReportServices(r *gin.RouterGroup, h *database.Handler) {
	// Create a new report handler
	reportHandler := services.NewReportHandler(h)

	// Set up routes
	r.POST("/party", reportHandler.CreatePartyReport)
	r.POST("/system", reportHandler.CreateSystemReport)
}
