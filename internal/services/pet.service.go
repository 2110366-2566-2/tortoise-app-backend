package services

import (
	"log"
	"net/http"

	"github.com/2110366-2566-2/tortoise-app-backend/internal/database"
	"github.com/2110366-2566-2/tortoise-app-backend/internal/models"
	"github.com/2110366-2566-2/tortoise-app-backend/internal/storage"
	"github.com/2110366-2566-2/tortoise-app-backend/pkg/utils"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PetHandler struct {
	dbHandler  *database.Handler
	stgHandler *storage.Handler
}

func NewPetHandler(db *database.Handler, stg *storage.Handler) *PetHandler {
	return &PetHandler{
		dbHandler:  db,
		stgHandler: stg,
	}
}

// func (h *PetHandler) GetAllPets(c *gin.Context) {
// 	pets, err := h.dbHandler.GetAllPetCards(c)
// 	if err != nil {
// 		log.Println("Error: ", err)
// 		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": "failed to get all pets"})
// 		return
// 	}
// 	c.JSON(http.StatusOK, gin.H{"success": true, "count": len(*pets), "data": &pets})
// }

func (h *PetHandler) GetPetBySeller(c *gin.Context) {
	// prevent xss attack
	sanitizedInput := utils.SanitizeString(c.Param("userID"))
	pets, err := h.dbHandler.GetPetBySeller(c, sanitizedInput)
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
// @Summary Get single pet by petID
// @tags pets
// @Description Get single pet by petID
// @id GetPetByPetID
// @produce json
// @Param petID path string true "ID of the pet to perform the operation on"
// @Router /api/v1/pets/pet/{petID} [get]
// @Success 200 {object} models.PetResponse
// @Failure 500 {object} models.ErrorResponse "internal server error"
func (h *PetHandler) GetPetByPetID(c *gin.Context) {
	id := c.Param("petID")
	sanitizedInput := utils.SanitizeString(id)
	pet, err := h.dbHandler.GetPetByPetID(c, sanitizedInput)
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
// @Summary Get filtered pets
// @tags pets
// @Description Get filtered pets by filter params
// @id GetFilteredPets
// @produce json
// @Router /api/v1/pets/ [get]
// @Param category query string false "Category of pet"
// @Param species query string false "Species of pet"
// @Param sex query string false "Sex of pet"
// @Param behavior query string false "Behavior of pet"
// @Param minAge query int false "Minimum age of pet"
// @Param maxAge query int false "Maximum age of pet"
// @Param minWeight query int false "Minimum weight of pet"
// @Param maxWeight query int false "Maximum weight of pet"
// @Param minPrice query int false "Minimum price of pet"
// @Param maxPrice query int false "Maximum price of pet"
// @Success 200 {object} models.PetCardResponse
// @Failure 400 {object} models.ErrorResponse "bad request"
// @Failure 500 {object} models.ErrorResponse "internal server error"
func (h *PetHandler) GetFilteredPets(c *gin.Context) {

	// prevent http parameter pollution

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

	// Validate parameters
	category, err = utils.ValidateArrayQueryParam(category)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "invalid query parameter"})
		return
	}
	species, err = utils.ValidateArrayQueryParam(species)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "invalid query parameter"})
		return
	}
	sex, err = utils.ValidateArrayQueryParam(sex)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "invalid query parameter"})
		return
	}
	behavior, err = utils.ValidateArrayQueryParam(behavior)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "invalid query parameter"})
		return
	}

	// Validate and convert query parameters to integers
	minAge, err = utils.GetIntQueryParam(minAgeStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "invalid minAge value"})
		return
	}
	maxAge, err = utils.GetIntQueryParam(maxAgeStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "invalid maxAge value"})
		return
	}
	minWeight, err = utils.GetIntQueryParam(minWeightStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "invalid minWeight value"})
		return
	}
	maxWeight, err = utils.GetIntQueryParam(maxWeightStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "invalid maxWeight value"})
		return
	}
	minPrice, err = utils.GetIntQueryParam(minPriceStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "invalid minPrice value"})
		return
	}
	maxPrice, err = utils.GetIntQueryParam(maxPriceStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "invalid maxPrice value"})
		return
	}

	pets, err := h.dbHandler.GetFilteredPetCards(c, category, species, sex, behavior, minAge, maxAge, minWeight, maxWeight, minPrice, maxPrice)
	if err != nil {
		log.Println("Error: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": "failed to get filtered pets"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "count": len(*pets), "data": &pets})
}

func (h *PetHandler) CreatePet(c *gin.Context) {
	var pet models.Pet
	if err := c.BindJSON(&pet); err != nil {
		log.Println("Error: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "failed to bind pet"})
		return
	}

	// prevent xss attack
	utils.PetSanitize(&pet)

	// check if seller exists
	sellerID := utils.SanitizeString(c.Param("userID"))

	_, err := h.dbHandler.GetSellerBySellerID(c, sellerID)
	if err != nil {
		log.Println("Error: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "seller not found"})
		return
	}

	pet.ID = primitive.NewObjectID()

	// convert string to objID
	userObjID, err := primitive.ObjectIDFromHex(c.Param("userID"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		return
	}
	pet.Seller_id = userObjID

	// convert media from base64 to url
	if len(pet.Media) > 0 {
		folder := "pets/" + pet.Seller_id.Hex()
		urls, err := h.stgHandler.AddImage(c, pet.ID.Hex(), folder, pet.Media)
		if err != nil {
			if err.Error() == "invalid base64 image string" {
				c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "invalid base64 image string"})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": "failed to upload media"})
			return
		}
		pet.Media = urls
	}

	// create pet
	err = h.dbHandler.CreateOnePet(c, pet.Seller_id, &pet)
	if err != nil {
		log.Println("Error: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"success": true, "data": &pet})
}

// UpdatePet godoc
// @Summary Update pet
// @Description Update pet by pet ID
// @Tags pets
// @Accept json
// @Produce json
// @Param petID path string true "ID of the pet to perform the operation on"
// @Param Pet body models.PetRequest true "Pet object that needs to be updated"
// @Router /api/v1/pets/pet/{petID} [put]
// @Success 200 {object} models.PetResponse "return updated pet"
// @Failure 400 {object} models.ErrorResponse "bad request"
// @Failure 500 {object} models.ErrorResponse "internal server error"
func (h *PetHandler) UpdatePet(c *gin.Context) {
	var data bson.M
	c.BindJSON(&data)

	// prevent xss attack
	utils.BsonSanitize(&data)

	// find if have media to upload
	if media, ok := data["media"]; ok {
		urls, err := h.stgHandler.AddImage(c, c.Param("petID"), "pets", media.(string))
		if err != nil {
			if err.Error() == "invalid base64 image string" {
				c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "invalid base64 image string"})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": "failed to upload media"})
			return
		}
		data["media"] = urls
	}

	petId := utils.SanitizeString(c.Param("petID"))

	res, err := h.dbHandler.UpdateOnePet(c, petId, data)
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
// @Summary Delete pet
// @Description Delete pet by pet ID and delete pet from user's pets
// @Tags pets
// @Accept json
// @Produce json
// @Param petID path string true "ID of the pet to delete"
// @Router /api/v1/pets/pet/{petID} [delete]
// @Success 200 {object} models.DeletePetResponse "return deleted count"
// @Failure 500 {object} models.ErrorResponse "internal server error"
func (h *PetHandler) DeletePet(c *gin.Context) {
	// get pet
	petID := utils.SanitizeString(c.Param("petID"))
	pet, err := h.dbHandler.GetPetByPetID(c, petID)
	if err != nil {
		log.Println("Error: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": "pet not found"})
		return
	}
	// delete pet
	res, err := h.dbHandler.DeleteOnePet(c, petID)
	if err != nil {
		log.Println("Error: ", err)
		errorMsg := "failed to delete pet"
		if err.Error()[:18] == "failed to get pet:" {
			errorMsg = "pet not found"
		}
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": errorMsg})
		return
	}

	// if pet has media, delete media
	if len(pet.Media) > 0 {
		// delete media
		folder := "pets/" + pet.Seller_id.Hex()
		if err := h.stgHandler.DeleteImage(c, petID, folder); err != nil {
			log.Println("Error: ", err)
			c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": "failed to delete media"})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "deletedCount": res.DeletedCount})
}

func (h *PetHandler) GetMasterData(c *gin.Context) {
	masterData, err := h.dbHandler.GetAllMasterData(c)
	if err != nil {
		c.JSON(500, gin.H{"success": false, "error": "failed to get master data"})
		return
	}
	c.JSON(200, gin.H{"success": true, "count": len(*masterData), "data": masterData})
}

func (h *PetHandler) GetMasterDataByCategory(c *gin.Context) {
	catagory := utils.SanitizeString(c.Param("category"))
	masterData, err := h.dbHandler.GetMasterDataByCategory(c, catagory)
	if err != nil {
		log.Println("Error: ", err)
		c.JSON(500, gin.H{"success": false, "error": "failed to get master data by category"})
		return
	}
	c.JSON(200, gin.H{"success": true, "data": masterData})
}

func (h *PetHandler) GetCategories(c *gin.Context) {
	categories, err := h.dbHandler.GetCategories(c)
	if err != nil {
		log.Println("Error: ", err)
		c.JSON(500, gin.H{"success": false, "error": "failed to get categories"})
		return
	}
	c.JSON(200, gin.H{"success": true, "count": len(categories.Categories), "data": categories})
}
