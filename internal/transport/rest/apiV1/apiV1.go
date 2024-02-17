package apiV1

import (
	"net/http"

	"github.com/2110366-2566-2/tortoise-app-backend/internal/database"
	"github.com/2110366-2566-2/tortoise-app-backend/internal/services"
	"github.com/gin-gonic/gin"
	"github.com/2110366-2566-2/tortoise-app-backend/configs"
	"github.com/auth0/go-jwt-middleware"
	// "github.com/dgrijalva/jwt-go"
	"github.com/form3tech-oss/jwt-go"
)

func PetController(r *gin.RouterGroup, h *database.Handler) {
	// Create a new pet handler
	petHandler := services.NewPetHandler(h)

	// Set up routes
	r.GET("/", petHandler.GetAllPets)
	r.GET("/filter", petHandler.GetFilteredPets)
	r.GET("/:petID", petHandler.GetPetByPetID)
	r.GET("/seller/:userID", petHandler.GetPetBySeller)

	r.POST("/seller/:userID", petHandler.CreatePet)
	r.PUT("/:petID", petHandler.UpdatePet)
	r.DELETE("/:petID", petHandler.DeletePet)
}

func UserServices(r *gin.RouterGroup, h *database.Handler) {
	// Set up routes
	r.POST("/login", func(c *gin.Context) {
		services.LoginHandler(c, h)
	})

	r.POST("/register", func(c *gin.Context) {
		services.RegisterHandler(c, h)
	})
}

func SetupRoutes(r *gin.Engine, h *database.Handler) {
	env, err := configs.LoadConfig()
	if err != nil {
		panic(err)
	}

	// Set up routes
	apiV1 := r.Group("/api/v1")

	apiV1.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "This is API v1.0.0")
	})

	
	UserServices(apiV1.Group("/user"), h)

	// Add JWT middleware to check the token
	apiV1.Use(jwtMiddleware(env))

	PetController(apiV1.Group("/pets"), h)
}

func jwtMiddleware(env configs.EnvVars) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the JWT secret key from the environment
		secretKey := env.JWT_SECRET

		// Create a new JWT middleware
		authMiddleware := jwtmiddleware.New(jwtmiddleware.Options{
			ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
				return []byte(secretKey), nil
			},
			SigningMethod: jwt.SigningMethodHS256,
		})

		// Handle the JWT middleware
		err := authMiddleware.CheckJWT(c.Writer, c.Request)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}

		c.Next()
	}
}
