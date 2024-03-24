package apiV1

import (
	"github.com/2110366-2566-2/tortoise-app-backend/internal/database"
	"github.com/2110366-2566-2/tortoise-app-backend/internal/services"
	"github.com/gin-gonic/gin"
)

func UnauthorizedRoutes(r *gin.RouterGroup, h *database.Handler) {
	// login and register
	r.POST("/login", func(c *gin.Context) {
		services.LoginHandler(c, h)
	})

	r.POST("/register", func(c *gin.Context) {
		services.RegisterHandler(c, h)
	})

	// user services without token
	userHandler := services.NewUserHandler(h)
	user := r.Group("/user")

	user.POST("/sentotp", userHandler.SentOTP)
	user.POST("/checkotp", userHandler.ValidateOTP)

	user.POST("/recoverusername", userHandler.RecoveryUsername)
	user.POST("/checkvalidemail", userHandler.CheckMail)

	// admin services without token
	r.POST("/admin/login", func(c *gin.Context) {
		services.LoginHandlerForAdmin(c, h)
	})
}
