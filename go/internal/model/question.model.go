package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Question struct {
	ID           primitive.ObjectID `bson:"_id" json:"_id"`
	Title        string             `bson:"title,omitempty" json:"title,omitempty"`
	Description  string             `bson:"description,omitempty" json:"description,omitempty"`
	DictionaryID primitive.ObjectID `bson:"dictionary_id" json:"dictionary_id"`
	Dictionaries []Dictionary       `bson:"dictionaries,omitempty" json:"dictionaries,omitempty"`
	Point        int64              `bson:"point,omitempty" json:"point,omitempty"`
	Option       []Option           `bson:"options,omitempty" json:"options,omitempty"`
}
