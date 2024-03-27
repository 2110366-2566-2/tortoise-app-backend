package apiV1

import (
	"github.com/2110366-2566-2/tortoise-app-backend/internal/database"
	"github.com/2110366-2566-2/tortoise-app-backend/internal/services"
	"github.com/gin-gonic/gin"
)

func SellerServices(r *gin.RouterGroup, h *database.Handler) {
	// Create a new seller handler
	sellerHandler := services.NewSellerHandler(h)

	sellerAdmin := r.Group("/")
	admin := r.Group("/")

	sellerAdmin.Use(roleMiddleware("seller", "admin"))
	admin.Use(roleMiddleware("admin"))

	// Set up routes
	admin.GET("/", sellerHandler.GetAllSellers)
	sellerAdmin.GET("/:sellerID", sellerHandler.GetSeller)
}
