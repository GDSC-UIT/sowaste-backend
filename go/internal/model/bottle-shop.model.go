package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Location struct {
	Latitude       float32 `bson:"latitude,omitempty" json:"latitude,omitempty"`
	Longtitude     float32 `bson:"longtitude,omitempty" json:"longtitude,omitempty"`
	LocationString string  `bson:"location_string,omitempty" json:"location_string,omitempty"`
	LocationMap    string  `bson:"location_map,omitempty" json:"location_map,omitempty"`
	Locale         string  `bson:"locale,omitempty" json:"locale,omitempty"`
}

type BottleShop struct {
	ID       primitive.ObjectID `bson:"_id" json:"_id"`
	Location Location           `bson:"location,omitempty" json:"location,omitempty"`
	Status   string             `bson:"status,omitempty" json:"status,omitempty"`
	Rate     uint8              `bson:"rate,omitempty" json:"rate,omitempty"`
	Contact  string             `bson:"contact,omitempty" json:"contact,omitempty"`
	Images   []string           `bson:"images,omitempty" json:"images,omitempty"`
}
