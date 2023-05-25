package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Badge struct {
	ID           primitive.ObjectID `bson:"_id" json:"_id"`
	DisplayImage string             `bson:"display_image" json:"display_image"`
	Condition    string             `bson:"condition" json:"condition"`
}
