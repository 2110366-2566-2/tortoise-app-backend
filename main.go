package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	mongoURI, dbName, port string
)

func main() {
	// Load .env file
	err := godotenv.Load(".env")
	if err != nil {
		panic("Error loading .env file")
	}

	mongoURI = os.Getenv("MONGO_URI")
	dbName = os.Getenv("DB_NAME")
	port = os.Getenv("PORT")

	// Connect to the database with connection pooling
	clientOpts := options.Client().SetMaxPoolSize(10) // Set an appropriate pool size
	fmt.Println("this is ", mongoURI)
	client, err := mongo.Connect(context.Background(), clientOpts.ApplyURI(mongoURI))
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := client.Disconnect(context.Background()); err != nil {
			panic(err)
		}
	}()

	// Create handler with database
	handler := newHandler(client.Database(dbName))

	// Set up routes
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "The server is running.")
	})
	r.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, "Ready to dev PetPal App !!!")
	})
	r.GET("/users", handler.usersHandler)
	// r.GET("/hospitals", handler.hospitalsHandler)

	// Run server
	r.Run(":" + port)
}

type User struct {
	ID           primitive.ObjectID `json:"id"`
	Name         string             `json:"name"`
	Gender       string             `json:"gender"`
	Phone_number string             `json:"phone_number"`
	Image        string             `json:"image"`
	Role         int32              `json:"role"`
	Email        string             `json:"email"`
	Password     string             `json:"password"`
	Address      struct {
		Province     string `json:"province"`
		District     string `json:"district"`
		Sub_district string `json:"sub_district"`
		Postal_code  string `json:"postal_code"`
		Street       string `json:"street"`
		Building     string `json:"building"`
		House_number string `json:"house_number"`
	} `json:"address"`
}

type Handler struct {
	db *mongo.Database
}

func newHandler(db *mongo.Database) *Handler {
	return &Handler{db}
}

func (h *Handler) usersHandler(c *gin.Context) {
	var users []User

	// Use context timeout for database queries
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Access the "users" collection
	collection := h.db.Collection("users")

	// Use projection to fetch all rows
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	// Decode results into users slice
	if err := cursor.All(ctx, &users); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	// Return users in the response
	c.JSON(http.StatusOK, &users)
}
