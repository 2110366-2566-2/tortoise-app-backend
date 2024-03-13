package database

import (
	"context"
	// "fmt"

	"github.com/2110366-2566-2/tortoise-app-backend/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetTransactionByID(ctx context.Context, h *Handler, userID primitive.ObjectID, role string) ([]*models.TransactionWithDetails, error) {

	var transactions []*models.TransactionWithDetails
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
		transactionWithDetails := models.TransactionWithDetails{Transaction: transaction}
		transactions = append(transactions, &transactionWithDetails)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return transactions, nil
}
