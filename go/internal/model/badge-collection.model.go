package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type BadgeCollection struct {
	ID      primitive.ObjectID `bson:"_id" json:"_id"`
	Badge   []Badge            `bson:"badge,omitempty" json:"badge,omitempty"`
	BadgeID primitive.ObjectID `bson:"badge_id" json:"badge_id"`
	UserID  string             `bson:"uid" json:"uid"`
}
