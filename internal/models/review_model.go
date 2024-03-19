package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Review struct {
	ID              primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Reviewer_id     primitive.ObjectID `json:"reviewer_id,omitempty" bson:"reviewer_id,omitempty"`
	Reviewee_id     primitive.ObjectID `json:"reviewee_id,omitempty" bson:"reviewee_id,omitempty"`
	Rating_score    float32            `json:"rating_score" bson:"rating_score"`
	Description     string             `json:"description" bson:"description"`
	Comment_records []Comments         `json:"comment_records" bson:"comment_records"`
}

type Comments struct {
	User_id primitive.ObjectID `json:"user_id" bson:"user_id"`
	Comment string             `json:"comment" bson:"comment"`
}
