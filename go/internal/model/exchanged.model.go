package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Exchanged struct {
	ID       primitive.ObjectID `bson:"_id" json:"_id"`
	UserID   primitive.ObjectID `bson:"user_id,omitempty" json:"user_id,omitempty"`
	RewardID primitive.ObjectID `bson:"reward_id,omitempty" json:"reward_id,omitempty"`
}
