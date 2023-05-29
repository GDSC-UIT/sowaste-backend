package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type QuizResult struct {
	ID           primitive.ObjectID `bson:"_id" json:"_id"`
	Total        int                `bson:"total,omitempty" json:"total,omitempty"`
	DictionaryID primitive.ObjectID `bson:"dictionary_id" json:"dictionary_id"`
	UserID       string             `bson:"uid" json:"uid"`
	Dictionary   []Dictionary       `bson:"dictionary,omitempty" json:"dictionary,omitempty"`
}
