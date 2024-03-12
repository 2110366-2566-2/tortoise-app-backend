package database

import (
	"context"
	// "fmt"

	"github.com/2110366-2566-2/tortoise-app-backend/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	// "go.mongodb.org/mongo-driver/mongo"
)

func GetTransactionByID(ctx context.Context, h *Handler, userID primitive.ObjectID, role string) ([]*models.Transaction, error) {

	var transactions []*models.Transaction
	var filter bson.M

	if role == "seller" {
		filter = bson.M{"seller_id": userID}
	} else {
		filter = bson.M{"buyer_id": userID}
	}

	cursor, err := h.db.Collection("transactions").Find(ctx, filter)
    if err != nil {
        return nil, err
    }
    defer cursor.Close(ctx)

    for cursor.Next(ctx) {
		var transaction models.Transaction
        if err := cursor.Decode(&transaction); err != nil {
            return nil, err
        }
        transactions = append(transactions, &transaction)
    }

    if err := cursor.Err(); err != nil {
        return nil, err
    }

    return transactions, nil
}
