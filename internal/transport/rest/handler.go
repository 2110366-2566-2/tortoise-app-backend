package rest

import (
	"net/http"

	"github.com/2110366-2566-2/tortoise-app-backend/internal/database"
	"github.com/2110366-2566-2/tortoise-app-backend/internal/services"
	"github.com/gin-gonic/gin"
)

func TestHandler(c *gin.Context) {
	c.JSON(http.StatusOK, "Ready to dev PetPal App !!!")
}

func RootHandler(c *gin.Context) {
	c.JSON(http.StatusOK, "The server is running.")
}

func PetController(r *gin.RouterGroup, h *database.Handler) {
	// Create a new pet handler
	petHandler := services.NewPetHandler(h)

	// Set up routes
	r.GET("/", petHandler.GetAllPets)
	r.GET("/:petID", petHandler.GetPetByPetID)
	r.GET("/seller/:userID", petHandler.GetPetBySeller)
	r.POST("/:userID", petHandler.CreatePet)

	// Print mock user id for testing
	// fmt.Printf("\nMock user id: %s\n\n", user.ID.Hex())
}

func SetupRoutes(r *gin.Engine, h *database.Handler) {
	r.GET("/test", TestHandler)
	r.GET("/", RootHandler)
	PetController(r.Group("/pets"), h)
}
