package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Dictionary struct {
	ID               primitive.ObjectID `bson:"_id" json:"_id"`
	Name             string             `bson:"name" json:"name"`
	Type             string             `bson:"type" json:"type"`
	ShortDescription string             `bson:"short_description" json:"short_description"`
	Description      string             `bson:"description" json:"description"`
	Uri              string             `bson:"uri" json:"uri"`
	Lessons          []Lesson           `bson:"lessons" json:"lessons"`
	Quizzes          []Quiz             `bson:"quizzes" json:"quizzes"`
	DisplayImage     string             `bson:"display_image" json:"display_image"`
}
