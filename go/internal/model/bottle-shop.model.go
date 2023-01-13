package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Location struct {
	X              int32  `bson:"x" json:"x"`
	Y              int32  `bson:"y" json:"y"`
	LocationString string `bson:"location_string" json:"location_string"`
}

type BottleShop struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Location Location           `bson:"location" json:"location"`
	Status   string             `bson:"status,omitempty" json:"status,omitempty"`
	Point    int64              `bson:"point,omitempty" json:"point,omitempty"`
}
