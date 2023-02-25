package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Dictionary struct {
	ID               primitive.ObjectID `bson:"_id" json:"_id"`
	Name             string             `bson:"name,omitempty" json:"name,omitempty"`
	IsOrganic        bool               `bson:"is_organic" json:"is_organic"`
	Recyable         bool               `bson:"recyable" json:"recyable"`
	ShortDescription string             `bson:"short_description,omitempty" json:"short_description,omitempty"`
	Description      string             `bson:"description,omitempty" json:"description,omitempty"`
	Uri              string             `bson:"uri,omitempty" json:"uri,omitempty"`
	Questions        []Question         `bson:"questions,omitempty" json:"questions,omitempty"`
	DisplayImage     string             `bson:"display_image,omitempty" json:"display_image,omitempty"`
}
