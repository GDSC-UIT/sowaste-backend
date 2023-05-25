package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Credential struct {
	AccessToken string `bson:"access_token,omitempty" json:"access_token,omitempty"`
	Password    string `bson:"password,omitempty" json:"password,omitempty"`
}

type User struct {
	ID           primitive.ObjectID `bson:"_id" json:"_id"`
	UID          string             `bson:"uid,omitempty" json:"uid,omitempty"`
	FullName     string             `bson:"full_name,omitempty" json:"full_name,omitempty"`
	Email        string             `bson:"email,omitempty" json:"email,omitempty"` // primary key
	Credential   Credential         `bson:"credential,omitempty" json:"credential,omitempty"`
	RewardPoint  int                `bson:"reward_point,omitempty" json:"reward_point,omitempty"`
	DisplayImage string             `bson:"display_image,omitempty" json:"display_image,omitempty"`
}
