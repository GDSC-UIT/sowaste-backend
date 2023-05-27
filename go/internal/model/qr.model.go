package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type QR struct {
	ID       primitive.ObjectID `bson:"_id" json:"_id"`
	UserIDs  []string           `bson:"uids" json:"uids"`
	Point    int                `bson:"point" json:"point"`
	ExpireAt primitive.DateTime `bson:"expire_at" json:"expire_at"`
	IssuedAt primitive.DateTime `bson:"issued_at" json:"issued_at"`
}
