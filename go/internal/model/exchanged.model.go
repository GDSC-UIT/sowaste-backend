package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Exchanged struct {
	ID       primitive.ObjectID `bson:"_id" json:"_id"`
	UserID   string             `bson:"uid" json:"-"`
	RewardID primitive.ObjectID `bson:"reward_id" json:"-"`
	Reward   []Reward           `bson:"reward,omitempty" json:"reward,omitempty"`
}
