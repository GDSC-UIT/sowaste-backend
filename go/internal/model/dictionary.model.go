package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Dictionary struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Category    string             `bson:"category,omitempty" json:"category,omitempty"`
	Type        string             `bson:"type,omitempty" json:"type,omitempty"`
	Description string             `bson:"description,omitempty" json:"description,omitempty"`
	Uri         string             `bson:"uri,omitempty" json:"uri,omitempty"`
}
