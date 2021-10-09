package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"Appointy-Instagram/data"
)

type UserHandler struct {
	userCollection *mongo.Collection
}

func NewUserHandler(col *mongo.Collection) *UserHandler {
	return &UserHandler{
		userCollection: col,
	}
}

func (h *UserHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.getUser(w, r)
	case http.MethodPost:
		h.createUser(w, r)
	default:
		http.Error(w, "Not allowed", http.StatusMethodNotAllowed)
	}
}

func (h *UserHandler) createUser(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	fmt.Println(body)

	user := &data.User{}

	err = nil
	err = json.Unmarshal(body, user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	fmt.Println(user)

	userResult, err1 := h.userCollection.InsertOne(context.Background(), user)
	if err1 != nil {
		fmt.Print(err1)
	}
	fmt.Println("Success: ", userResult)

	w.Write([]byte("still works"))
}

func (h *UserHandler) getUser(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Path[len("/users/"):]
	fmt.Println(id)

	user := &data.User{}

	userResult := h.userCollection.FindOne(context.Background(), bson.D{{"_id", id}})
	err := userResult.Decode(user)
	fmt.Println("User result: \n", user)

	if err != nil {
		w.Write([]byte("unable to get data"))
	} else {
		userJson, err := json.Marshal(user)
		if err != nil {
			w.Write([]byte(fmt.Sprintf("Unable to marshal JSON due to err: \v", err)))
		} else {
			w.Write(userJson)
		}

	}

}
