package apiV1

import (
	"net/http"

	"github.com/2110366-2566-2/tortoise-app-backend/internal/database"
	"github.com/2110366-2566-2/tortoise-app-backend/internal/services"
	"github.com/gin-gonic/gin"
)

func PetController(r *gin.RouterGroup, h *database.Handler) {
	// Create a new pet handler
	petHandler := services.NewPetHandler(h)

	// Set up routes
	r.GET("/", petHandler.GetAllPets)
	r.GET("/:petID", petHandler.GetPetByPetID)
	r.GET("/seller/:userID", petHandler.GetPetBySeller)
	r.POST("/seller/:userID", petHandler.CreatePet)
	r.PUT("/:petID", petHandler.UpdatePet)
	r.DELETE("/:petID", petHandler.DeletePet)
}

func UserServices(r *gin.RouterGroup, h *database.Handler) {
	// Set up routes
	r.POST("/login", func(c *gin.Context) {
		services.LoginHandler(c, h)
	})

	r.POST("/register", func(c *gin.Context) {
		services.RegisterHandler(c, h)
	})
}

func SetupRoutes(r *gin.Engine, h *database.Handler) {
	apiV1 := r.Group("/api/v1")

	apiV1.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "This is API v1.0.0")
	})

	PetController(apiV1.Group("/pets"), h)

	UserServices(apiV1.Group("/user"), h)
}
