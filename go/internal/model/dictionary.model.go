package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Dictionary struct {
	ID   primitive.ObjectID `bson:"_id" json:"_id"`
	Name string             `bson:"name" json:"name"`
	// IsOrganic          bool               `bson:"is_organic" json:"is_organic"` // deprecated
	// Recyable           bool               `bson:"recyable" json:"recyable"`
	// ShortDescription   string             `bson:"short_description" json:"short_description"`
	// Description        string             `bson:"description" json:"description"`
	// Uri                string             `bson:"uri" json:"uri"`
	// Questions          []Question         `bson:"questions" json:"questions"`
	DisplayImage       string   `bson:"display_image" json:"display_image"`
	Types              []string `bson:"types" json:"types"`
	GoodToKnow         string   `bson:"good_to_know" json:"good_to_know"`
	RecyclableItems    []string `bson:"recyclable_items" json:"recyclable_items"`
	NonRecyclableItems []string `bson:"non_recyclable_items" json:"non_recyclable_items"`
	HowToRecyclable    string   `bson:"how_to_recyclable" json:"how_to_recyclable"`
}
