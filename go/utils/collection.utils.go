package utils

import (
	"os"

	"go.mongodb.org/mongo-driver/mongo"
)

func GetDatabaseCollection(name string, db *mongo.Client) *mongo.Collection {
	return db.Database(os.Getenv("DB_DATABASE")).Collection(name)
}

func ContainsElement(arr []string, target string) bool {
	for _, element := range arr {
		if element == target {
			return true
		}
	}
	return false
}
