package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type PartyReport struct {
	ID              primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Reporter_id     primitive.ObjectID `json:"reporter_id,omitempty" bson:"reporter_id,omitempty"`
	Reportee_id     primitive.ObjectID `json:"reportee_id,omitempty" bson:"reportee_id,omitempty"`
	Description     string             `json:"description" bson:"description"`
	Is_solved       bool               `json:"is_solved" bson:"is_solved"`
}

type SystemReport struct {
	ID              primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Reporter_id     primitive.ObjectID `json:"reporter_id,omitempty" bson:"reporter_id,omitempty"`
	Description     string             `json:"description" bson:"description"`
	Is_solved       bool               `json:"is_solved" bson:"is_solved"`
}
