package database

import (
	"context"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Client struct {
	Db *mongo.Client
}

var client Client

func (db *Client) ConnectDb() {

	var err error

	client.Db, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(os.Getenv("MONGO_URI")))

	if err != nil {
		panic(err)
	}

	if err := client.Db.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected and pinged.")

}

func (db *Client) DisconnetDb() {
	var err error

	if err = client.Db.Disconnect(context.TODO()); err != nil {
		panic(err)
	}
}
