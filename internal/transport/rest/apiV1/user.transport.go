package apiV1

import (
	"github.com/2110366-2566-2/tortoise-app-backend/internal/database"
	"github.com/2110366-2566-2/tortoise-app-backend/internal/services"
	"github.com/gin-gonic/gin"
)

func UserServices(r *gin.RouterGroup, h *database.Handler) {

	// Create a new user handler
	userHandler := services.NewUserHandler(h)

	r.GET("/:userID", userHandler.GetUserByUserID)
	r.PUT("/passwd/:userID", userHandler.UpdateUserPasswd)
	r.PUT("/:userID", userHandler.UpdateUser)
	r.DELETE("/:userID", userHandler.DeleteUser)
	// r.GET("/token/session", func(c *gin.Context) {
	//     services.GetSessionToken(c, h)
	// })

	r.POST("/forgotpasswd", userHandler.UpdateForgotPassword)

	// Get me
	r.GET("/me", userHandler.WhoAmI)

}