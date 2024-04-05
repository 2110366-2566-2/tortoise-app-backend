package database

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Handler struct {
	db *mongo.Database
}

func NewHandler(db *mongo.Database) *Handler {
	return &Handler{db: db}
}

func ConnectMongo(uri, dbName string, timeout time.Duration) (*mongo.Database, context.CancelFunc, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	// defer cancel()

	clientOpts := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(ctx, clientOpts)
	if err != nil {
		return nil, cancel, err
	}

	log.Println("Connected to MongoDB")

	return client.Database(dbName), cancel, nil
}

func CloseMongo(db *mongo.Database, cancel context.CancelFunc) error {
	// cancel the context
	cancel()
	// if err return err
	return db.Client().Disconnect(context.Background())
}
