package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type DIY struct {
	ID               primitive.ObjectID `bson:"_id" json:"_id"`
	Title            string             `bson:"title,omitempty" json:"title,omitempty"`
	Description      string             `bson:"description,omitempty" json:"description,omitempty"`
	ShortDescription string             `bson:"short_description,omitempty" json:"short_description,omitempty"`
	Source           string             `bson:"source,omitempty" json:"source,omitempty"`
	CreatedAt        primitive.DateTime `bson:"created_at,omitempty" json:"created_at,omitempty"`
	DisplayImage     string             `bson:"display_image,omitempty" json:"display_image,omitempty"`
	DictionaryID     primitive.ObjectID `bson:"dictionary_id" json:"dictionary_id"`
	Dictionary       []Dictionary       `bson:"dictionary,omitempty" json:"dictionary,omitempty"`
}
