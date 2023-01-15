package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Dictionary struct {
	ID          primitive.ObjectID `bson:"_id" json:"_id"`
	Category    string             `bson:"category" json:"category"`
	Type        string             `bson:"type" json:"type"`
	Description string             `bson:"description" json:"description"`
	Uri         string             `bson:"uri" json:"uri"`
	Lessons     []Lesson           `bson:"lessons,omitempty" json:"lessons,omitempty"`
	Quizzes     []Quiz             `bson:"quizzes,omitempty" json:"quizzes,omitempty"`
}
