package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Reward struct {
	ID           primitive.ObjectID `bson:"_id" json:"_id"`
	DisplayImage string             `bson:"display_image" json:"display_image"`
	Point        int                `bson:"point" json:"point"`
}
