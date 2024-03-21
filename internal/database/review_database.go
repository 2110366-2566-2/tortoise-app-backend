package database

import (
	"context"
	"fmt"
	"math"

	"github.com/2110366-2566-2/tortoise-app-backend/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (h *Handler) CreateReview(ctx context.Context, review *models.Review) (*mongo.InsertOneResult, error) {
	// Insert a new transaction
	review.ID = primitive.NewObjectID()

	// check if rating score is within 0-5
	if review.Rating_score < 0 {
		review.Rating_score = 0
	}
	if review.Rating_score > 5 {
		review.Rating_score = 5
	}

	// round to 2 decimal places
	review.Rating_score = math.Round(review.Rating_score*100) / 100

	review.Comment_records = make([]models.Comments, 0)
	res, err := h.db.Collection("reviews").InsertOne(ctx, review)
	if err != nil {
		return nil, fmt.Errorf("failed to create review")
	}
	return res, nil
}

func (h *Handler) GetReviewByUserID(ctx context.Context, UserID string) (*[]models.Review, error) {
	userObjID, err := primitive.ObjectIDFromHex(UserID)
	if err != nil {
		return nil, fmt.Errorf("failed to convert reviewID to ObjectID: %v", err)
	}
	filter := bson.M{"reviewee_id": userObjID}
	cursor, err := h.db.Collection("reviews").Find(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("failed to find reviews: %v", err)
	}
	defer cursor.Close(ctx)

	var reviews []models.Review
	for cursor.Next(ctx) {
		var review models.Review
		if err := cursor.Decode(&review); err != nil {
			return nil, fmt.Errorf("failed to decode document: %v", err)
		}
		reviews = append(reviews, review)
	}
	return &reviews, nil
}

func (h *Handler) GetReviewByReviewID(ctx context.Context, reviewID string) (*models.Review, error) {
	reviewObjID, err := primitive.ObjectIDFromHex(reviewID)
	if err != nil {
		return nil, fmt.Errorf("failed to convert reviewID to ObjectID: %v", err)
	}
	filter := bson.M{"_id": reviewObjID}
	var review models.Review
	if err := h.db.Collection("reviews").FindOne(ctx, filter).Decode(&review); err != nil {
		return nil, fmt.Errorf("failed to find review: %v", err)
	}
	return &review, nil
}

func (h *Handler) CreateComment(ctx context.Context, reviewID string, comment models.Comments) (*models.Review, error) {
	// convert string to objID
	// var updateDoc bson.D
	// for k, v := range data {
	// 	updateDoc = append(updateDoc, bson.E{Key: k, Value: v})
	// }

	reviewObjID, err := primitive.ObjectIDFromHex(reviewID)
	if err != nil {
		return nil, fmt.Errorf("failed to convert reviewID to ObjectID: %v", err)
	}

	filter := bson.M{"_id": reviewObjID}
	update := bson.M{"$push": bson.D{{Key: "comment_records", Value: comment}}}
	_, err = h.db.Collection("reviews").UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, fmt.Errorf("failed to update comment's review: %v", err)
	}

	var review models.Review
	err = h.db.Collection("reviews").FindOne(ctx, filter).Decode(&review)
	if err != nil {
		return nil, fmt.Errorf("failed to find review: %v", err)
	}
	return &review, nil
}

func (h *Handler) DeleteReview(ctx context.Context, reviewID, role string, userID primitive.ObjectID) (*bson.M, error) {
	reviewObjID, err := primitive.ObjectIDFromHex(reviewID)
	if err != nil {
		return nil, fmt.Errorf("failed to convert to ObjectID")
	}
	var review models.Review
	if err := h.db.Collection("reviews").FindOne(ctx, bson.M{"_id": reviewObjID}).Decode(&review); err != nil {
		return nil, fmt.Errorf("failed to find review: %v", err)
	}
	commentCount := len(review.Comment_records)

	if review.Reviewer_id != userID && role != "admin" {
		return nil, fmt.Errorf("unauthorized")
	}
	filter := bson.M{"_id": reviewObjID}
	_, err = h.db.Collection("reviews").DeleteOne(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("failed to delete review: %v", err)
	}
	res := bson.M{"comments_deleted": commentCount, "review_id": review.ID.Hex()}
	return &res, nil
}
