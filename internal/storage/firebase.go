package storage

import (
	"context"
	"log"

	"cloud.google.com/go/storage"
	"google.golang.org/api/option"
)

type Handler struct {
	stg *storage.Client
}

func NewHandler(stg *storage.Client) *Handler {
	return &Handler{stg: stg}
}

func ConnectFirebase(ctx context.Context, serviceAccountKeyPath string) (*storage.Client, error) {
	opt := option.WithCredentialsFile(serviceAccountKeyPath)
	app, err := storage.NewClient(ctx, opt)
	if err != nil {
		return nil, err
	}

	log.Println("Connected to Firebase")

	return app, nil
}
