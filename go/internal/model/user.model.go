package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Credential struct {
	AccessToken string `bson:"access_token,omitempty" json:"access_token,omitempty"`
}

type User struct {
	ID         primitive.ObjectID `bson:"_id" json:"_id"`
	FullName   string             `bson:"full_name,omitempty" json:"full_name,omitempty"`
	Email      string             `bson:"email,omitempty" json:"email,omitempty"` // primary key
	Password   string             `bson:"password,omitempty" json:"password,omitempty"`
	Credential Credential         `bson:"credential,omitempty" json:"credential,omitempty"`
}
