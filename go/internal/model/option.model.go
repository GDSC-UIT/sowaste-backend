package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Option struct {
	ID           primitive.ObjectID `bson:"_id" json:"_id"`
	Description  string             `bson:"description" json:"description"` //* question *//
	QuestionID   primitive.ObjectID `bson:"question_id" json:"question_id"`
	DictionaryID primitive.ObjectID `bson:"dictionary_id" json:"dictionary_id"`
	IsCorrect    bool               `bson:"is_correct" json:"is_correct"`
}
