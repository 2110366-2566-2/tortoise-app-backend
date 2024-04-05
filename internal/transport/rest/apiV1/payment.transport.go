package apiV1

import (
	"github.com/2110366-2566-2/tortoise-app-backend/configs"
	"github.com/2110366-2566-2/tortoise-app-backend/internal/database"
	"github.com/2110366-2566-2/tortoise-app-backend/internal/services"
	"github.com/gin-gonic/gin"
)

func PaymentServices(r *gin.RouterGroup, h *database.Handler, env configs.EnvVars) {
	// Create a new buyer handler
	buyerHandler := services.NewPaymentHandler(h, env)

	// Set up routes
	r.POST("/create", buyerHandler.CreatePayment)
	r.POST("/confirm", buyerHandler.ConfirmPayment)
}
