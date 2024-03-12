package database

import (
	"context"
	"fmt"

	"github.com/2110366-2566-2/tortoise-app-backend/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (h *Handler) CreateTransaction(ctx context.Context, transaction *models.Transaction) (*mongo.InsertOneResult, error) {
	// Insert a new transaction
	transaction.ID = primitive.NewObjectID()
	res, err := h.db.Collection("transactions").InsertOne(ctx, transaction)
	if err != nil {
		return nil, fmt.Errorf("failed to create transaction")
	}
	return res, nil
}

func (h *Handler) GetTransactionByTransactionID(ctx context.Context, transactionID string) (*models.Transaction, error) {
	var transaction models.Transaction
	transactionObjID, err := primitive.ObjectIDFromHex(transactionID)
	if err != nil {
		return nil, fmt.Errorf("failed to convert transactionID to ObjectID: %v", err)
	}
	filter := bson.M{"_id": transactionObjID}
	err = h.db.Collection("transactions").FindOne(ctx, filter).Decode(&transaction)
	if err != nil {
		return nil, fmt.Errorf("failed to find transaction: %v", err)
	}
	return &transaction, nil
}

func (h *Handler) GetTransactionByUserID(ctx context.Context, userID string) (*[]models.Transaction, error) {
	var transactions []models.Transaction

	// get role
	role, err := h.GetUserRole(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to get user role: %v", err)
	}

	var filter bson.M
	if role == "seller" {
		filter = bson.M{"seller_id": userID}
	} else if role == "buyer" {
		filter = bson.M{"buyer_id": userID}
	}
	cursor, err := h.db.Collection("transactions").Find(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("failed to find transactions: %v", err)
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var transaction models.Transaction
		if err := cursor.Decode(&transaction); err != nil {
			return nil, fmt.Errorf("failed to decode transaction: %v", err)
		}
		transactions = append(transactions, transaction)
	}

	return &transactions, nil
}
