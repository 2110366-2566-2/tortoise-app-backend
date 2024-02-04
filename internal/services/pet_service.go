package services

import (
	"fmt"
	"net/http"

	"github.com/2110366-2566-2/tortoise-app-backend/internal/database"
	"github.com/2110366-2566-2/tortoise-app-backend/internal/models"
	"github.com/gin-gonic/gin"
)

type PetHandler struct {
	handler *database.Handler
}

func NewPetHandler(handler *database.Handler) *PetHandler {
	return &PetHandler{handler: handler}
}

// GetAllPets godoc
// @Summary Get all pets
// @Description Get all pets collection
// @Endpoint /pets
func (h *PetHandler) GetAllPets(c *gin.Context) {
	pets, err := h.handler.GetAllPets(c)
	if err != nil {
		fmt.Println("Error: ", err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, &pets)
}

// GetPetBySeller godoc
// @Summary Get pets by seller
// @Description Get pets by seller id
// @Endpoint /pets/seller/:userID
func (h *PetHandler) GetPetBySeller(c *gin.Context) {
	userID := c.Param("userID")
	pets, err := h.handler.GetPetBySeller(c, userID)
	if err != nil {
		fmt.Println("Error: ", err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, &pets)
}

// GetPetByPetID godoc
// @Summary Get pet by pet id
// @Description Get pet by pet id
// @Endpoint /pets/:petID
func (h *PetHandler) GetPetByPetID(c *gin.Context) {
	id := c.Param("petID")
	pet, err := h.handler.GetPetByPetID(c, id)
	if err != nil {
		fmt.Println("Error: ", err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, &pet)
}

// CreatePet godoc
// @Summary Create pet
// @Description Create pet with user id
// @Endpoint /pets/:userID
func (h *PetHandler) CreatePet(c *gin.Context) {
	userID := c.Param("userID")
	var pet models.Pet
	c.BindJSON(&pet)
	res, err := h.handler.CreateOnePet(c, userID, &pet)
	if err != nil {
		fmt.Println("Error: ", err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusCreated, &res.InsertedID)
}

// UpdatePet godoc
// @Summary Update pet
// @Description Update pet by pet id
// @Endpoint /pets/:petID
func (h *PetHandler) UpdatePet(c *gin.Context) {
	// id := c.Param("petID")
	// var pet models.Pet
	// c.BindJSON(&pet)
	// for i, p := range test.Pets {
	// 	if p.ID.Hex() == id {
	// 		test.Pets[i] = pet
	// 		c.JSON(http.StatusOK, gin.H{
	// 			"message": "Pet updated successfully",
	// 			"petID":   id,
	// 		})
	// 		return
	// 	}
	// }
	// c.JSON(http.StatusNotFound, "Pet not found")
	c.JSON(http.StatusNotImplemented, "Not implemented")
}

// DeletePet godoc
// @Summary Delete pet
// @Description Delete pet by pet id
// @Endpoint /pets/:petID
func (h *PetHandler) DeletePet(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, "Not implemented")
}
