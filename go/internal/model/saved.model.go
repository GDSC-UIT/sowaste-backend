package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Saved struct {
	ID           primitive.ObjectID `bson:"_id" json:"_id"`
	UserID       string             `bson:"uid" json:"uid"`
	DictionaryID primitive.ObjectID `bson:"dictionary_id" json:"dictionary_id"`
}
