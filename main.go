package main

import (
	"fmt"
	"log"
	"net/http"

	"Appointy-Instagram/data"
	"Appointy-Instagram/handlers"
)

func main() {
	// Initializing the mongodb client
	mongoClient := data.ConnectToDb()

	// connecting to collections
	userCollection := mongoClient.Database("Insta").Collection("Users")
	postCollection := mongoClient.Database("Insta").Collection("Posts")

	//Initializing habdlers from each route
	userHandler := handlers.NewUserHandler(userCollection)
	postHandler := handlers.NewPostHandler(postCollection)
	postUserHandler := handlers.NewPostUserHandler(postCollection)

	//Routing the server using different handlers
	http.Handle("/users/", userHandler)
	http.Handle("/posts/", postHandler)
	http.Handle("/posts/users/", postUserHandler)

	// Starting the server at localhost:8080
	fmt.Println("Starting the server at localhost:8080 ...")
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
