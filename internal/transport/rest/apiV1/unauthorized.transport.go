package apiV1

import (
	"github.com/2110366-2566-2/tortoise-app-backend/internal/database"
	"github.com/2110366-2566-2/tortoise-app-backend/internal/services"
	"github.com/2110366-2566-2/tortoise-app-backend/internal/storage"
	"github.com/gin-gonic/gin"
)

func UnauthorizedRoutes(r *gin.RouterGroup, h *database.Handler, stg *storage.Handler) {
	// login and register
	r.POST("/login", func(c *gin.Context) {
		services.LoginHandler(c, h)
	})

	r.POST("/register", func(c *gin.Context) {
		services.RegisterHandler(c, h, stg)
	})

	// user services without token
	userHandler := services.NewUserHandler(h, stg)
	user := r.Group("/user")

	user.POST("/sentotp", userHandler.SentOTP)
	user.POST("/checkotp", userHandler.ValidateOTP)

	user.POST("/recoverusername", userHandler.RecoveryUsername)
	user.POST("/checkvalidemail", userHandler.CheckMail)

	// admin services without token
	admin := r.Group("/admin")
	admin.POST("/login", func(c *gin.Context) {
		services.LoginHandlerForAdmin(c, h)
	})
	admin.POST("/register", func(c *gin.Context) {
		services.AdminRegisterHandler(c, h, stg)
	})
}
