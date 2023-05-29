package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Badge struct {
	ID           primitive.ObjectID `bson:"_id" json:"_id"`
	DisplayImage string             `bson:"display_image,omitempty" json:"display_image,omitempty"`
	Condition    string             `bson:"condition,omitempty" json:"condition,omitempty"`
}
