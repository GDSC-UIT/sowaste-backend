package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Saved struct {
	ID           primitive.ObjectID `bson:"_id" json:"_id"`
	UserID       primitive.ObjectID `bson:"user_id" json:"user_id"`
	DictionaryID primitive.ObjectID `bson:"dictionary_id" json:"dictionary_id"`
}
