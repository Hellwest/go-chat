package database

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client = getClient()

func getClient() *mongo.Client {
	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb://admin:chat@localhost:27017/admin")

	// Connect to MongoDB
	fmt.Println("Connecting to the database")
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		panic(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		panic(err)
	}

	fmt.Println("Connected to MongoDB")

	return client
}

func GetCollection(collectionName string) *mongo.Collection {
	return Client.Database("chat").Collection(collectionName)
}
