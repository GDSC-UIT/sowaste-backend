package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Category struct {
	ID           primitive.ObjectID `bson:"_id" json:"_id"`
	Name         string             `bson:"name" json:"name"`
	DisplayImage string             `bson:"display_image" json:"display_image"`
	DictonaryID  primitive.ObjectID `bson:"dictionary_id" json:"dictionary_id"`
}
