package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Quiz struct {
	ID           primitive.ObjectID `bson:"_id" json:"_id"`
	Title        string             `bson:"title" json:"title"`
	Description  string             `bson:"description" json:"description"`
	DictionaryID primitive.ObjectID `bson:"dictionary_id" json:"dictionary_id"`
	Dictionaries []Dictionary       `bson:"dictionaries" json:"dictionaries"`
	Point        int64              `bson:"point" json:"point"`
	Question     []Question         `bson:"questions" json:"questions"`
}
