package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Quiz struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	DictionaryID primitive.ObjectID `bson:"dictionary_id,omitempty" json:"-,omitempty"`
	Dictionary   []Dictionary       `bson:"dictionaries,omitempty" json:"dictionaries,omitempty"`
	Status       string             `bson:"status,omitempty" json:"status,omitempty"`
	Point        int64              `bson:"point,omitempty" json:"point,omitempty"`
}
