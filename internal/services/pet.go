package services

import (
	"fmt"
	"net/http"

	"github.com/2110366-2566-2/tortoise-app-backend/internal/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// =========== mock data ===========

var pets []models.Pet

var user = models.User{
	ID:           primitive.NewObjectID(),
	Name:         "John",
	Surname:      "Doe",
	Gender:       "M",
	Phone_number: "0812345678",
	Image:        "https://www.google.com",
	Role:         1,
	Email:        "test@test.com",
	Password:     "123456",
	Address: struct {
		Province     string `bson:"province"`
		District     string `bson:"district"`
		Sub_district string `bson:"sub_district"`
		Postal_code  string `bson:"postal_code"`
		Street       string `bson:"street"`
		Building     string `bson:"building"`
		House_number string `bson:"house_number"`
	}{Province: "Bangkok",
		District:     "Bangkok",
		Sub_district: "Bangkok",
		Postal_code:  "10100",
		Street:       "Bangkok",
		Building:     "Bangkok",
		House_number: "Bangkok",
	},
	Pets: make([]string, 0),
}

var users []models.User = []models.User{user}

// =========== mock data ===========

// GetAllPets godoc
// @Summary Get all pets
// @Description Get all pets collection
// @Endpoint /pets
func getAllPets(c *gin.Context) {
	c.JSON(http.StatusOK, &pets)
}

// GetPetBySeller godoc
// @Summary Get pets by seller
// @Description Get pets by seller id
// @Endpoint /pets/seller/:userID
func getPetBySeller(c *gin.Context) {
	userID := c.Param("userID")
	var petsByUserID []models.Pet
	for _, user := range users {
		if user.ID.Hex() == userID {
			for _, pet := range pets {
				if pet.Seller_id == user.ID.Hex() {
					petsByUserID = append(petsByUserID, pet)
				}
			}
			c.JSON(http.StatusOK, &petsByUserID)
			return
		}
		c.JSON(http.StatusNotFound, "Pet not found")
	}
}

// GetPetByPetID godoc
// @Summary Get pet by pet id
// @Description Get pet by pet id
// @Endpoint /pets/:petID
func getPetByPetID(c *gin.Context) {
	id := c.Param("petID")
	for _, pet := range pets {
		if pet.ID.Hex() == id {
			c.JSON(http.StatusOK, &pet)
			return
		}
	}
	c.JSON(http.StatusNotFound, "Pet not found")
}

// CreatePet godoc
// @Summary Create pet
// @Description Create pet with user id
// @Endpoint /pets/:userID
func createPet(c *gin.Context) {
	userID := c.Param("userID")
	var pet models.Pet
	petID := primitive.NewObjectID()
	c.BindJSON(&pet)
	pet.ID = petID
	pet.Seller_id = userID
	for i, user := range users {
		if user.ID.Hex() == userID {
			users[i].Pets = append(users[i].Pets, pet.ID.Hex())
			pets = append(pets, pet)
			c.JSON(http.StatusOK, gin.H{
				"message":  "Pet created successfully",
				"petID":    pet.ID.Hex(),
				"sellerID": userID,
			})
			return
		}
	}
	c.JSON(http.StatusNotFound, "Seller not found")
}

func PetController(r *gin.RouterGroup) {
	r.GET("/", getAllPets)
	r.GET("/:petID", getPetByPetID)
	r.GET("/seller/:userID", getPetBySeller)
	r.POST("/:userID", createPet)

	// Print mock user id for testing
	fmt.Printf("\nMock user id: %s\n\n", user.ID.Hex())
}
