package database

import (
	"context"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Instance struct {
	Source *mongo.Client
}

var Client Instance

func (db *Instance) ConnectDb() {
	var err error

	Client.Source, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(os.Getenv("MONGO_URI")))

	if err != nil {
		panic(err)
	}

	if err := Client.Source.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected and pinged.")
}

func (db *Instance) DisconnetDb() {
	var err error

	if err = Client.Source.Disconnect(context.TODO()); err != nil {
		panic(err)
	}
}
