package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Credential struct {
	AccessToken string `bson:"access_token" json:"access_token"`
	Password    string `bson:"password" json:"password"`
}

type User struct {
	ID           primitive.ObjectID `bson:"_id" json:"_id"`
	UID          string             `bson:"uid" json:"uid"`
	FullName     string             `bson:"full_name" json:"full_name"`
	Email        string             `bson:"email" json:"email"` // primary key
	Credential   Credential         `bson:"credential" json:"credential"`
	RewardPoint  int                `bson:"reward_point" json:"reward_point"`
	DisplayImage string             `bson:"display_image" json:"display_image"`
}
