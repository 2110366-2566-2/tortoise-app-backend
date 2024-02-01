package app

import (
	"github.com/2110366-2566-2/tortoise-app-backend/internal/transport/rest"
	"github.com/gin-gonic/gin"
)

func Run() {
	r := gin.Default()

	rest.SetupRoutes(r)

	r.Run(":8080")
}
