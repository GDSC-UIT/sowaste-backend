package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Location struct {
	Latitude       float32 `bson:"latitude" json:"latitude"`
	Longtitude     float32 `bson:"longtitude" json:"longtitude"`
	LocationString string  `bson:"location_string" json:"location_string"`
	LocationMap    string  `bson:"location_map" json:"location_map"`
	Locale         string  `bson:"locale" json:"locale"`
}

type BottleShop struct {
	ID       primitive.ObjectID `bson:"_id" json:"_id"`
	Location Location           `bson:"location" json:"location"`
	Status   string             `bson:"status" json:"status"`
	Rate     uint8              `bson:"rate" json:"rate"`
	Contact  string             `bson:"contact" json:"contact"`
	Images   []string           `bson:"images" json:"images"`
}
