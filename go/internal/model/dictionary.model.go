package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Dictionary struct {
	ID               primitive.ObjectID `bson:"_id" json:"_id"`
	Name             string             `bson:"name" json:"name"`
	IsOrganic        bool               `bson:"is_organic" json:"is_organic"`
	Recyable         bool               `bson:"recyable" json:"recyable"`
	ShortDescription string             `bson:"short_description" json:"short_description"`
	Description      string             `bson:"description" json:"description"`
	Uri              string             `bson:"uri" json:"uri"`
	Lessons          []Lesson           `bson:"lessons" json:"lessons"`
	Questions        []Question         `bson:"questions" json:"questions"`
	DisplayImage     string             `bson:"display_image" json:"display_image"`
}
