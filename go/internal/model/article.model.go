package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Article struct {
	ID               primitive.ObjectID `bson:"_id" json:"_id"`
	Title            string             `bson:"title" json:"title"`
	Description      string             `bson:"description" json:"description"`
	ShortDescription string             `bson:"short_description" json:"short_description"`
	Source           string             `bson:"source" json:"source"`
	CreatedAt        primitive.DateTime `bson:"created_at" json:"created_at"`
}
