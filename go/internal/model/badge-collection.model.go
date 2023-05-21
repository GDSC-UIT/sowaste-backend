package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type BadgeCollection struct {
	ID      primitive.ObjectID `bson:"_id" json:"_id"`
	BadgeID primitive.ObjectID `bson:"badge_id,omitempty" json:"badge_id,omitempty"`
	UserID  primitive.ObjectID `bson:"user_id,omitempty" json:"user_id,omitempty"`
}
