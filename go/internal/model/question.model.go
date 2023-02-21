package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Question struct {
	ID           primitive.ObjectID `bson:"_id" json:"_id"`
	Title        string             `bson:"title" json:"title"`
	DictionaryID primitive.ObjectID `bson:"dictionary_id" json:"dictionary_id"`
	Dictionaries []Dictionary       `bson:"dictionaries" json:"dictionaries"`
	Point        int64              `bson:"point" json:"point"`
	Option       []Option           `bson:"options" json:"options"`
}
