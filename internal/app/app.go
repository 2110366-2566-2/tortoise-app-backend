package app

import (
	"github.com/Armph/tortoise-app-backend/internal/transport/rest"
	"github.com/gin-gonic/gin"
	"github.com/taufanmahaputra/forex/internal/transport/rest"
)

func Run() {
	r := gin.Default()

	rest.SetupRoutes(r)

	r.Run(":8080")
}
