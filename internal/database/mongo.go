package database

import (
	"context"
	"log"
	"time"

	"github.com/2110366-2566-2/tortoise-app-backend/internal/models"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Handler struct {
	db *mongo.Database
}

func NewHandler(db *mongo.Database) *Handler {
	return &Handler{db: db}
}

// func (h *Handler) CreatePet(ctx context.Context, pet models.Pet) error {
// 	_, err := h.db.Collection("pets").InsertOne(ctx, pet)
// 	return err
// }

// func (h *Handler) GetPetByID(ctx context.Context, id string) (models.Pet, error) {
// 	var pet models.Pet
// 	err := h.db.Collection("pets").FindOne(context.Background(), map[string]string{"_id": id}).Decode(&pet)
// 	return pet, err
// }

func (h *Handler) GetAllPets(ctx context.Context) (*[]models.Pet, error) {
	var pets []models.Pet
	// check data base is connected
	if err := h.db.Client().Ping(ctx, nil); err != nil {
		return nil, err
	}
	cursor, err := h.db.Collection("pets").Find(ctx, map[string]string{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var pet models.Pet
		if err := cursor.Decode(&pet); err != nil {
			return nil, err
		}
		pets = append(pets, pet)
	}
	return &pets, nil
}

func ConnectMongo(uri, dbName string, timeout time.Duration) (*mongo.Database, error) {
	ctx, _ := context.WithTimeout(context.Background(), timeout)
	// defer cancel()

	clientOpts := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(ctx, clientOpts)
	if err != nil {
		return nil, err
	}

	log.Println("Connected to MongoDB")

	return client.Database(dbName), nil
}

func CloseMongo(db *mongo.Database) error {
	// if err return err
	return db.Client().Disconnect(context.Background())
}
