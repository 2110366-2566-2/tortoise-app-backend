package apiV1

import (
	"github.com/2110366-2566-2/tortoise-app-backend/internal/database"
	"github.com/2110366-2566-2/tortoise-app-backend/internal/services"
	"github.com/2110366-2566-2/tortoise-app-backend/internal/storage"
	"github.com/gin-gonic/gin"
)

func PetController(r *gin.RouterGroup, db *database.Handler, stg *storage.Handler) {
	// Create a new pet handler
	petHandler := services.NewPetHandler(db, stg)

	allUser := r.Group("/")
	allUser.Use(roleMiddleware("seller", "admin", "buyer"))

	sellerAdmin := r.Group("/")
	sellerAdmin.Use(roleMiddleware("seller", "admin"))

	// Set up routes
	// allUser.GET("/old", petHandler.GetAllPets)
	allUser.GET("/", petHandler.GetFilteredPets)
	allUser.GET("/pet/:petID", petHandler.GetPetByPetID)
	allUser.GET("/seller/:userID", petHandler.GetPetBySeller)

	sellerAdmin.POST("/seller/:userID", petHandler.CreatePet)
	sellerAdmin.PUT("/pet/:petID", petHandler.UpdatePet)
	sellerAdmin.DELETE("/pet/:petID", petHandler.DeletePet)

	// master data
	allUser.GET("/master", petHandler.GetMasterData)
	allUser.GET("/master/:category", petHandler.GetMasterDataByCategory)
	allUser.GET("/master/categories", petHandler.GetCategories)
}
