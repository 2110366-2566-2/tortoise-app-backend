package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
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
	// r.Use("/pets", PetController)
}
