package services

import (
	"log"
	"net/http"

	"github.com/2110366-2566-2/tortoise-app-backend/internal/database"
	"github.com/2110366-2566-2/tortoise-app-backend/internal/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
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
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, &pets)
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
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, &pets)
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
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, &pet)
}

// CreatePet godoc
// @Method POST
// @Summary Create pet
// @Description Create pet with user id
// @Endpoint /api/v1/pets/:userID
func (h *PetHandler) CreatePet(c *gin.Context) {
	var pet models.Pet
	if err := c.BindJSON(&pet); err != nil {
		log.Println("Error: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	pet.Seller_id = c.Param("userID")
	res, err := h.handler.CreateOnePet(c, pet.Seller_id, &pet)
	if err != nil {
		log.Println("Error: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, res)
}

// UpdatePet godoc
// @Method PUT
// @Pa
// @Summary Update pet
// @Description Update pet by pet id
// @Endpoint /api/v1/pets/:petID
func (h *PetHandler) UpdatePet(c *gin.Context) {
	var data bson.M
	c.BindJSON(&data)
	res, err := h.handler.UpdateOnePet(c, c.Param("petID"), data)
	if err != nil {
		log.Println("Error: ", err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, &res)
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
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, &res)
}
