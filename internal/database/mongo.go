package database

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

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
