package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type QuizResult struct {
	ID           primitive.ObjectID `bson:"_id" json:"_id"`
	Total        int                `bson:"total" json:"total"`
	DictionaryID primitive.ObjectID `bson:"dictionary_id" json:"-"`
	UserID       string             `bson:"uid" json:"uid"`
	Dictionary   []Dictionary       `bson:"dictionary" json:"dictionary"`
	IsDone       bool               `bson:"is_done" json:"is_done"`
}
