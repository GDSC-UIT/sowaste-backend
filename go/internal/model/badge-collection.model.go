package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type BadgeCollection struct {
	ID      primitive.ObjectID `bson:"_id" json:"_id"`
	BadgeID primitive.ObjectID `bson:"badge_id,omitempty" json:"badge_id,omitempty"`
	UserID  string             `bson:"uid,omitempty" json:"uid,omitempty"`
}
