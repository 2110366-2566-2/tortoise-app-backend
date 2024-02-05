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
	r.POST("/:userID", petHandler.CreatePet)
	r.PUT("/:petID", petHandler.UpdatePet)
	r.DELETE("/:petID", petHandler.DeletePet)
}

func UserController(r *gin.RouterGroup, h *database.Handler) {
	// Create a new user handler
	userHandler := services.NewUserHandler(h)

	// Set up routes
	r.GET("/", userHandler.GetAllUsers)
	r.GET("/:userID", userHandler.GetUserByUserID)
}

func SetupRoutes(r *gin.Engine, h *database.Handler) {
	apiV1 := r.Group("/api/v1")

	apiV1.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "This is API v1.0.0")
	})

	apiV1.POST("/login", services.LoginHandler)
	PetController(apiV1.Group("/pets"), h)
	UserController(apiV1.Group("/users"), h)
}
