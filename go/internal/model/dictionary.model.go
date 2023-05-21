package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Dictionary struct {
	ID                 primitive.ObjectID `bson:"_id" json:"_id"`
	Name               string             `bson:"name,omitempty" json:"name,omitempty"`
	IsOrganic          bool               `bson:"is_organic" json:"is_organic"` // deprecated
	Recyable           bool               `bson:"recyable" json:"recyable"`
	ShortDescription   string             `bson:"short_description,omitempty" json:"short_description,omitempty"`
	Description        string             `bson:"description,omitempty" json:"description,omitempty"`
	Uri                string             `bson:"uri,omitempty" json:"uri,omitempty"`
	Questions          []Question         `bson:"questions,omitempty" json:"questions,omitempty"`
	DisplayImage       string             `bson:"display_image,omitempty" json:"display_image,omitempty"`
	Types              []string           `bson:"types,omitempty" json:"types,omitempty"`
	GoodToKnow         string             `bson:"good_to_know,omitempty" json:"good_to_know,omitempty"`
	RecyclableItems    []string           `bson:"recyable_items" json:"recyable_items"`
	NonRecyclableItems []string           `bson:"non_recyable_items" json:"non_recyable_items"`
	HowToRecyclable    string             `bson:"how_to_recycle,omitempty" json:"how_to_recycle,omitempty"`
}
