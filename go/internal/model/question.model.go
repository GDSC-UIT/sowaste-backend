package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Question struct {
	ID           primitive.ObjectID `bson:"_id" json:"_id"`
	Description  string             `bson:"description" json:"description"` //* question *//
	QuizID       primitive.ObjectID `bson:"quiz_id" json:"quiz_id"`
	DictionaryID primitive.ObjectID `bson:"dictionary_id" json:"dictionary_id"`
	IsCorrect    bool               `bson:"is_correct" json:"is_correct"`
}
