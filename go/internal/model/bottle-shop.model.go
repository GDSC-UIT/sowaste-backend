package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Location struct {
	X              float32 `bson:"x" json:"x"`
	Y              float32 `bson:"y" json:"y"`
	LocationString string  `bson:"location_string" json:"location_string"`
	LocationMap    string  `bson:"location_map" json:"location_map"`
}

type BottleShop struct {
	ID       primitive.ObjectID `bson:"_id" json:"_id"`
	Location Location           `bson:"location" json:"location"`
	Status   string             `bson:"status" json:"status"`
	Rate     uint8              `bson:"rate" json:"rate"`
}
