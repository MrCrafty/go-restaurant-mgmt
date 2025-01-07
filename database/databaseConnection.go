package database

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func DBInstance() *mongo.Client {
	MongoDb := "mongodb://admin:admin123@localhost:27017"
	fmt.Println(MongoDb)
	client, err := mongo.Connect(options.Client().ApplyURI(MongoDb))
	if err != nil {
		fmt.Println(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10)
	defer cancel()
	if err := client.Ping(ctx, nil); err != nil {
		fmt.Println(err)
	}
	fmt.Println("Connected to MongoDB!")
	return client
}

var Client = DBInstance()

func OpenCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	collection := client.Database("restaurant").Collection(collectionName)
	return collection
}
