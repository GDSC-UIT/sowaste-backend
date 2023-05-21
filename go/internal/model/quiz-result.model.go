package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type QuizResult struct {
	ID           primitive.ObjectID `bson:"_id" json:"_id"`
	Total        int32              `bson:"total,omitempty" json:"total,omitempty"`
	DictionaryID primitive.ObjectID `bson:"dictionary_id,omitempty" json:"dictionary_id,omitempty"`
	UserID       primitive.ObjectID `bson:"user_id,omitempty" json:"user_id,omitempty"`
}
