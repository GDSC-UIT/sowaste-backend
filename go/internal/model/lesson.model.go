package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Lesson struct {
	ID           primitive.ObjectID `bson:"_id" json:"_id"`
	Title        string             `bson:"title" json:"title"`
	Description  string             `bson:"description" json:"description"`
	DictionaryID primitive.ObjectID `bson:"dictionary_id" json:"dictionary_id"` //** Reference to dictionary **//
	Dictionaries []Dictionary       `bson:"dictionaries" json:"dictionaries"`
	Status       string             `bson:"status" json:"status"`
}
