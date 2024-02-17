package services

import (
	"log"
	"net/http"
	"strconv"

	"github.com/2110366-2566-2/tortoise-app-backend/internal/database"
	"github.com/2110366-2566-2/tortoise-app-backend/internal/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PetHandler struct {
	handler *database.Handler
}

func NewPetHandler(handler *database.Handler) *PetHandler {
	return &PetHandler{handler: handler}
}

// GetAllPets godoc
// @Method GET
// @Summary Get all pets
// @Description Get all pets cards
// @Endpoint /api/v1/pets
func (h *PetHandler) GetAllPets(c *gin.Context) {
	pets, err := h.handler.GetAllPetCards(c)
	if err != nil {
		log.Println("Error: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": "failed to get all pets"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "count": len(*pets), "data": &pets})
}

// GetPetBySeller godoc
// @Method GET
// @Summary Get pets by seller
// @Description Get pets by seller id
// @Endpoint /api/v1/pets/seller/:userID
func (h *PetHandler) GetPetBySeller(c *gin.Context) {
	pets, err := h.handler.GetPetBySeller(c, c.Param("userID"))
	if err != nil {
		log.Println("Error: ", err)
		errorMsg := "failed to get pets by seller"
		if err.Error() == "seller not found" {
			errorMsg = "seller not found"
		}
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": errorMsg})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "count": len(*pets), "data": &pets})
}

// GetPetByPetID godoc
// @Method GET
// @Summary Get pet by pet id
// @Description Get pet by pet id
// @Endpoint /api/v1/pets/:petID
func (h *PetHandler) GetPetByPetID(c *gin.Context) {
	id := c.Param("petID")
	pet, err := h.handler.GetPetByPetID(c, id)
	if err != nil {
		log.Println("Error: ", err)
		errorMsg := "failed to get pet by pet id"
		if err.Error() == "mongo: no documents in result" {
			errorMsg = "pet not found"
		}
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": errorMsg})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "data": &pet})
}

// GetPetFilteredPets godoc
// @Method GET
// @Summary Get filtered pets
// @Description Get filtered pets by filter params
// @Endpoint /api/v1/pets/filter
func (h *PetHandler) GetFilteredPets(c *gin.Context) {
	category := c.QueryArray("category")
	species := c.QueryArray("species")
	sex := c.QueryArray("sex")
	behavior := c.QueryArray("behavior")
	minAgeStr := c.Query("minAge")
	maxAgeStr := c.Query("maxAge")
	minWeightStr := c.Query("minWeight")
	maxWeightStr := c.Query("maxWeight")
	minPriceStr := c.Query("minPrice")
	maxPriceStr := c.Query("maxPrice")

	var minPrice, maxPrice, minAge, maxAge, minWeight, maxWeight int
	var err error

	if minPriceStr != "" {
		minPrice, err = strconv.Atoi(minPriceStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "invalid minPrice value"})
			return
		}
	}
	if maxPriceStr != "" {
		maxPrice, err = strconv.Atoi(maxPriceStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "invalid maxPrice value"})
			return
		}
	}
	if minAgeStr != "" {
		minAge, err = strconv.Atoi(minAgeStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "invalid minAge value"})
			return
		}
	}
	if maxAgeStr != "" {
		maxAge, err = strconv.Atoi(maxAgeStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "invalid maxAge value"})
			return
		}
	}
	if minWeightStr != "" {
		minWeight, err = strconv.Atoi(minWeightStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "invalid minWeight value"})
			return
		}
	}
	if maxWeightStr != "" {
		maxWeight, err = strconv.Atoi(maxWeightStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "invalid maxWeight value"})
			return
		}
	}

	pets, err := h.handler.GetFilteredPetCards(c, category, species, sex, behavior, minAge, maxAge, minWeight, maxWeight, minPrice, maxPrice)
	if err != nil {
		log.Println("Error: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": "failed to get filtered pets"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "count": len(*pets), "data": &pets})
}

// CreatePet godoc
// @Method POST
// @Summary Create pet
// @Description Create pet with user id
// @Endpoint /api/v1/pets/seller/:userID
func (h *PetHandler) CreatePet(c *gin.Context) {
	var pet models.Pet
	if err := c.BindJSON(&pet); err != nil {
		log.Println("Error: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "failed to bind pet"})
		return
	}
	// create new pet id
	pet.ID = primitive.NewObjectID()

	// convert string to objID
	petObjID, err := primitive.ObjectIDFromHex(c.Param("userID"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		return
	}
	pet.Seller_id = petObjID
	err = h.handler.CreateOnePet(c, pet.Seller_id, &pet)
	if err != nil {
		log.Println("Error: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"success": true, "data": &pet})
}

// UpdatePet godoc
// @Method PUT
// @Summary Update pet
// @Description Update pet by pet id
// @Endpoint /api/v1/pets/:petID
func (h *PetHandler) UpdatePet(c *gin.Context) {
	var data bson.M
	c.BindJSON(&data)
	res, err := h.handler.UpdateOnePet(c, c.Param("petID"), data)
	if err != nil {
		log.Println("Error: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": "failed to update pet"})
		return
	}
	// log result
	var updatedPet models.Pet
	err = res.Decode(&updatedPet)
	if err != nil {
		log.Println("Error decoding updated pet:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": "failed to decode updated pet"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "data": updatedPet})
}

// DeletePet godoc
// @Method DELETE
// @Summary Delete pet
// @Description Delete pet by pet id and delete pet from user's pets
// @Endpoint /api/v1/pets/:petID
func (h *PetHandler) DeletePet(c *gin.Context) {
	res, err := h.handler.DeleteOnePet(c, c.Param("petID"))
	if err != nil {
		log.Println("Error: ", err)
		errorMsg := "failed to delete pet"
		if err.Error()[:18] == "failed to get pet:" {
			errorMsg = "pet not found"
		}
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": errorMsg})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "deletedCount": res.DeletedCount})
}
