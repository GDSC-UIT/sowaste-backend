package services

import (
	"context"
	"fmt"

	"github.com/GDSC-UIT/sowaste-backend/go/internal/config"
	"github.com/GDSC-UIT/sowaste-backend/go/internal/database"
	"github.com/GDSC-UIT/sowaste-backend/go/internal/model"
	"go.mongodb.org/mongo-driver/bson"
)

func ReadManyDictionaries() []model.Dictionary {
	var Dictionaries []model.Dictionary

	//** Populate Dictionary field **//
	var aggPopulate = bson.M{"$lookup": bson.M{
		"from":         "Dictionary",    // the collection name
		"localField":   "dictionary_id", // the field on the child struct
		"foreignField": "_id",           // the field on the parent struct
		"as":           "dictionaries",  // the field to populate into
	}}
	collection := database.Client.Source.Database(config.GetDBConfig().DbName)
	cursor, _ := collection.Aggregate(context.TODO(), []bson.M{
		aggPopulate,
	})

	if err := cursor.All(context.TODO(), &Dictionaries); err != nil {
		fmt.Println("Error!")
		return nil
	}

	return Dictionaries
}
