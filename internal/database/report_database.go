package database

import (
	"context"
	"fmt"

	"github.com/2110366-2566-2/tortoise-app-backend/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)


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

func (h *Handler) GetReport(ctx context.Context, category string, is_solved *bool) (*[]models.PartyReport, *[]models.SystemReport, error) {
    var partyReport []models.PartyReport
    var systemReport []models.SystemReport
    var filter bson.M

    if is_solved != nil {
        filter = bson.M{"is_solved": *is_solved}
    }

    if category == "party" || category == "all" || category == "" {
        cursor, err := h.db.Collection("party_reports").Find(ctx, filter)
        if err != nil {
            return nil, nil, fmt.Errorf("failed to find party reports: %v", err)
        }
        defer cursor.Close(ctx)
        for cursor.Next(ctx) {
            var r models.PartyReport
            if err := cursor.Decode(&r); err != nil {
                return nil, nil, fmt.Errorf("failed to decode document: %v", err)
            }
            partyReport = append(partyReport, r)
        }
    }

    if category == "system" || category == "all" || category == "" {
        cursor, err := h.db.Collection("system_reports").Find(ctx, filter)
        if err != nil {
            return nil, nil, fmt.Errorf("failed to find system reports: %v", err)
        }
        defer cursor.Close(ctx)
        for cursor.Next(ctx) {
            var r models.SystemReport
            if err := cursor.Decode(&r); err != nil {
                return nil, nil, fmt.Errorf("failed to decode document: %v", err)
            }
            systemReport = append(systemReport, r)
        }
    }

    if category != "party" && category != "system" && category != "all" && category != "" {
        return nil, nil, fmt.Errorf("incorrect category: %v", category)
    }

    return &partyReport, &systemReport, nil
}
