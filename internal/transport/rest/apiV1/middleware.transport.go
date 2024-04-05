package apiV1

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/2110366-2566-2/tortoise-app-backend/configs"
	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	"github.com/form3tech-oss/jwt-go"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

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

		// Extract the token from the request
		tokenString := c.GetHeader("Authorization")
		tokenString = strings.TrimPrefix(tokenString, "Bearer ")

		// Parse the token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(secretKey), nil
		})

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}

		// Extract the role from the token
		claims := token.Claims.(jwt.MapClaims)
		userID, _ := primitive.ObjectIDFromHex(claims["userID"].(string))
		username := claims["username"].(string)
		role := claims["role"].(string)

		// Pass the role to the next middleware/handler
		c.Set("role", role)
		c.Set("userID", userID)
		c.Set("username", username)

		c.Next()
	}
}

func roleMiddleware(allowedRoles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		role, exists := c.Get("role")
		if !exists {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}

		for _, allowedRole := range allowedRoles {
			if role == allowedRole {
				c.Next()
				return
			}
		}

		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Forbidden"})
	}
}
