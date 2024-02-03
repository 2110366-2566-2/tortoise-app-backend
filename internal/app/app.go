package app

import (
	"github.com/2110366-2566-2/tortoise-app-backend/internal/transport/rest"
	"github.com/2110366-2566-2/tortoise-app-backend/internal/transport/rest/apiV1"
	"github.com/gin-gonic/gin"
)

func Run() {
	r := gin.Default()

	rest.SetupRoutes(r)

	apiV1.SetupRoutes(r)

	r.Run(":8080")
}
