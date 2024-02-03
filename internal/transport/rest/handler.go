package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/2110366-2566-2/tortoise-app-backend/internal/services"
)

func TestHandler(c *gin.Context) {
	c.JSON(http.StatusOK, "Ready to dev PetPal App !!!")
}

func RootHandler(c *gin.Context) {
	c.JSON(http.StatusOK, "The server is running.")
}

func SetupRoutes(r *gin.Engine) {
	r.GET("/test", TestHandler)
	r.GET("/", RootHandler)

	services.PetController(r.Group("/pets"))
}
