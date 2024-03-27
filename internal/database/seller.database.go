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

func (h *Handler) AddBankAccount(ctx context.Context, sellerID string, bankAccount models.BankAccount) (*mongo.UpdateResult, error) {
	// Convert sellerID to ObjectID
	sellerObjID, err := primitive.ObjectIDFromHex(sellerID)
	if err != nil {
		return nil, fmt.Errorf("failed to convert sellerID to ObjectID: %v", err)
	}

	// Update bank account field of the seller
	filter := bson.M{"_id": sellerObjID}
	update := bson.M{"$set": bson.M{"bank_account": bankAccount}}
	res, err := h.db.Collection("sellers").UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, fmt.Errorf("failed to update bank account: %v", err)
	}
	return res, nil
}

func (h *Handler) GetBankAccount(sellerID string) (*models.BankAccount, error) {
	// Convert sellerID to ObjectID
	sellerObjID, err := primitive.ObjectIDFromHex(sellerID)
	if err != nil {
		return nil, fmt.Errorf("failed to convert sellerID to ObjectID: %v", err)
	}

	// Get bank account of the seller
	filter := bson.M{"_id": sellerObjID}
	var seller models.Seller
	err = h.db.Collection("sellers").FindOne(context.Background(), filter).Decode(&seller)
	if err != nil {
		return nil, fmt.Errorf("failed to find seller: %v", err)
	}

	if seller.BankAccount == (models.BankAccount{}) {
		return nil, nil
	}

	return &seller.BankAccount, nil
}

func (h *Handler) DeleteBankAccount(sellerID string) (*mongo.UpdateResult, error) {
	// Convert sellerID to ObjectID
	sellerObjID, err := primitive.ObjectIDFromHex(sellerID)
	if err != nil {
		return nil, fmt.Errorf("failed to convert sellerID to ObjectID: %v", err)
	}

	// Delete bank account of the seller
	filter := bson.M{"_id": sellerObjID}
	update := bson.M{"$unset": bson.M{"bank_account": ""}}
	res, err := h.db.Collection("sellers").UpdateOne(context.Background(), filter, update)
	if err != nil {
		return nil, fmt.Errorf("failed to delete bank account: %v", err)
	}
	return res, nil
}

func (h *Handler) ChangeStatus(ctx context.Context, sellerID string, status string) (*mongo.UpdateResult, error) {
	// Convert sellerID to ObjectID
	sellerObjID, err := primitive.ObjectIDFromHex(sellerID)
	if err != nil {
		return nil, fmt.Errorf("failed to convert sellerID to ObjectID: %v", err)
	}

	// Update status of the seller
	filter := bson.M{"_id": sellerObjID}
	update := bson.M{"$set": bson.M{"status": status}}
	res, err := h.db.Collection("sellers").UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, fmt.Errorf("failed to update status: %v", err)
	}

	return res, nil
}

func (h *Handler) GetSellerBySellerID(ctx context.Context, sellerID string) (*models.Seller, error) {
	var seller models.Seller
	// Convert sellerID to ObjectID
	sellerObjID, err := primitive.ObjectIDFromHex(sellerID)
	if err != nil {
		return nil, fmt.Errorf("failed to convert sellerID to ObjectID: %v", err)
	}
	opts := options.FindOne().SetProjection(bson.M{"pets": 0})
	filter := bson.M{"_id": sellerObjID}
	err = h.db.Collection("sellers").FindOne(ctx, filter, opts).Decode(&seller)
	if err != nil {
		return nil, fmt.Errorf("failed to find seller: %v", err)
	}
	return &seller, nil
}

func (h *Handler) GetAllSellers(ctx context.Context, status string) (*[]models.Seller, error) {
	var sellers []models.Seller
	// Check if the status is valid
	if status != "verified" && status != "unverified" && status != "" {
		return nil, fmt.Errorf("invalid status")
	}
	var filter bson.M
	if status == "" {
		filter = bson.M{}
	} else {
		filter = bson.M{"status": status}
	}
	cursor, err := h.db.Collection("sellers").Find(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("failed to find sellers: %v", err)
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var seller models.Seller
		if err := cursor.Decode(&seller); err != nil {
			return nil, fmt.Errorf("failed to decode document: %v", err)
		}
		sellers = append(sellers, seller)
	}

	return &sellers, nil
}
