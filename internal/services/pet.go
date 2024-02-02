package services

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getAllPets(c *gin.Context) {
	c.JSON(http.StatusOK, "Get all pets")
}

func getPet(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, "Get pets"+id)
}

func PetController(r *gin.Engine) {
	r.GET("/", getAllPets)
	r.GET("/:id", getPet)
}
