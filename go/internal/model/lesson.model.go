package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Lesson struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	DictionaryID primitive.ObjectID `bson:"-" json:"dictionary_id,omitempty"`
	Dictionary   []Dictionary       `bson:"dictionaries,omitempty" json:"dictionaries,omitempty"`
	Status       string             `bson:"status,omitempty" json:"status,omitempty"`
	NextLesson   string             `bson:"next_lesson" json:"next_lesson"`
	PrevLesson   string             `bson:"prev_lesson" json:"prev_lesson"`
}
