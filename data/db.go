package data

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Function to connect to the mongodb server
func ConnectToDb() *mongo.Client {
	// Connection to db at localhost:27017
	mongoClient, err := mongo.Connect(context.Background(), &options.ClientOptions{
		Auth: &options.Credential{
			Username: "mongoadmin",
			Password: "secret",
		},
	})
	if err != nil {
		log.Fatalf("Unable to connect to db\n[Error]: %v", err)
	}

	//Creating user and post collections
	mongoClient.Database("Insta").CreateCollection(context.Background(), "Users")
	mongoClient.Database("Insta").CreateCollection(context.Background(), "Posts")

	return mongoClient
}
