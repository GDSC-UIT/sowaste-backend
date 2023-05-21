package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Category struct {
	ID           primitive.ObjectID `bson:"_id" json:"_id"`
	Name         string             `bson:"name,omitempty" json:"name,omitempty"`
	DisplayImage string             `bson:"display_image,omitempty" json:"display_image,omitempty"`
	DictonaryID  primitive.ObjectID `bson:"dictonary_id" json:"dictonary_id"`
}
