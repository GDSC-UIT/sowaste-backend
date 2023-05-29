package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Content struct {
	Title string   `bson:"title,omitempty" json:"title,omitempty"`
	Data  []string `bson:"data,omitempty" json:"data,omitempty"`
}

type Dictionary struct {
	ID               primitive.ObjectID `bson:"_id" json:"_id"`
	Name             string             `bson:"name" json:"name"`
	IsOrganic        bool               `bson:"is_organic,omitempty" json:"is_organic,omitempty"` // deprecated
	Recyable         bool               `bson:"recyable,omitempty" json:"recyable,omitempty"`
	ShortDescription string             `bson:"short_description,omitempty" json:"short_description,omitempty"`
	// Description        string             `bson:"description" json:"description"`
	// Uri                string             `bson:"uri" json:"uri"`
	DisplayImage       string   `bson:"display_image,omitempty" json:"display_image,omitempty"`
	Types              []string `bson:"types,omitempty" json:"types,omitempty"`
	GoodToKnow         string   `bson:"good_to_know,omitempty" json:"good_to_know,omitempty"`
	RecyclableItems    Content  `bson:"recyclable_items,omitempty" json:"recyclable_items,omitempty"`
	NonRecyclableItems Content  `bson:"non_recyclable_items,omitempty" json:"non_recyclable_items,omitempty"`
	HowToRecyclable    string   `bson:"how_to_recyclable,omitempty" json:"how_to_recyclable,omitempty"`
}
