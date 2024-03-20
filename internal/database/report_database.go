package database

import (
	"context"
	"fmt"

	"github.com/2110366-2566-2/tortoise-app-backend/internal/models"
	// "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// func (h *Handler) CreateReview(ctx context.Context, review *models.Review) (*mongo.InsertOneResult, error) {
// 	// Insert a new transaction
// 	review.ID = primitive.NewObjectID()
// 	review.Comment_records = make([]models.Comments, 0)
// 	res, err := h.db.Collection("reviews").InsertOne(ctx, review)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to create review")
// 	}
// 	return res, nil
// }

func (h *Handler) CreatePartyReport(ctx context.Context, report *models.PartyReport) (*mongo.InsertOneResult, error) {
	
	report.ID = primitive.NewObjectID()
	res, err := h.db.Collection("party_reports").InsertOne(ctx, report)
	if err != nil {
		return nil, fmt.Errorf("failed to create party report")
	}
	return res, nil
}

func (h *Handler) CreateSystemReport(ctx context.Context, report *models.SystemReport) (*mongo.InsertOneResult, error) {
	
	report.ID = primitive.NewObjectID()
	res, err := h.db.Collection("system_reports").InsertOne(ctx, report)
	if err != nil {
		return nil, fmt.Errorf("failed to create system report")
	}
	return res, nil
}
// func (h *Handler) GetReviewByUserID(ctx context.Context, UserID string) (*[]models.Review, error) {
// 	userObjID, err := primitive.ObjectIDFromHex(UserID)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to convert reviewID to ObjectID: %v", err)
// 	}
// 	filter := bson.M{"reviewee_id": userObjID}
// 	cursor, err := h.db.Collection("reviews").Find(ctx, filter)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to find reviews: %v", err)
// 	}
// 	defer cursor.Close(ctx)

// 	var reviews []models.Review
// 	for cursor.Next(ctx) {
// 		var review models.Review
// 		if err := cursor.Decode(&review); err != nil {
// 			return nil, fmt.Errorf("failed to decode document: %v", err)
// 		}
// 		reviews = append(reviews, review)
// 	}
// 	return &reviews, nil
// }

// func (h *Handler) CreateComment(ctx context.Context, reviewID string, data bson.M) (*models.Review, error) {
// 	// convert string to objID
// 	var updateDoc bson.D
// 	for k, v := range data {
// 		updateDoc = append(updateDoc, bson.E{Key: k, Value: v})
// 	}

// 	reviewObjID, err := primitive.ObjectIDFromHex(reviewID)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to convert reviewID to ObjectID: %v", err)
// 	}

// 	filter := bson.M{"_id": reviewObjID}
// 	update := bson.M{"$push": updateDoc}
// 	_, err = h.db.Collection("reviews").UpdateOne(ctx, filter, update)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to update comment's review: %v", err)
// 	}

// 	var review models.Review
// 	err = h.db.Collection("reviews").FindOne(ctx, filter).Decode(&review)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to find review: %v", err)
// 	}
// 	return &review, nil
// }
