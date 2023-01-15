package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Location struct {
	X              int32  `bson:"x" json:"x"`
	Y              int32  `bson:"y" json:"y"`
	LocationString string `bson:"location_string" json:"location_string"`
}

type BottleShop struct {
	ID       primitive.ObjectID `bson:"_id" json:"_id"`
	Location Location           `bson:"location" json:"location"`
	Status   string             `bson:"status" json:"status"`
	Point    int64              `bson:"point" json:"point"`
}
