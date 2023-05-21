package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Reward struct {
	ID           primitive.ObjectID `bson:"_id" json:"_id"`
	DisplayImage string             `bson:"display_image,omitempty" json:"display_image,omitempty"`
	Point        int8               `bson:"point,omitempty" json:"point,omitempty"`
}
