package database

import (
	"context"
	"fmt"

	"github.com/2110366-2566-2/tortoise-app-backend/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

func (h *Handler) UpdateTransaction(ctx context.Context, transactionID primitive.ObjectID, update bson.D) (*mongo.SingleResult, error) {
	// Put the transaction into the database
	res := h.db.Collection("transactions").FindOneAndUpdate(ctx, bson.M{"_id": transactionID}, update, options.FindOneAndUpdate().SetReturnDocument(options.After))
	if res.Err() != nil {
		return nil, fmt.Errorf("failed to update transaction")
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

func (h *Handler) GetTransactionByID(ctx context.Context, userID primitive.ObjectID, role string) ([]*models.TransactionWithDetails, error) {

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
