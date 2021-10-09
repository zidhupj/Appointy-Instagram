package handlers

import (
	"Appointy-Instagram/data"
	"Appointy-Instagram/functions"
	"context"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type PostHandler struct {
	postCollection *mongo.Collection
}

func NewPostHandler(col *mongo.Collection) *PostHandler {
	return &PostHandler{
		postCollection: col,
	}
}

func (h *PostHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.getPost(w, r)
	case http.MethodPost:
		h.createPost(w, r)
	default:
		http.Error(w, "Bad request", http.StatusBadRequest)
	}
}

func (h *PostHandler) createPost(w http.ResponseWriter, r *http.Request) {
	// getting post from request body
	post := &data.InPost{}
	ok := functions.ReadJson(w, r, post)
	if !ok {
		return
	}

	//validating post
	if err := functions.ValidatePost(post); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Generating random id for post
	rand.Seed(time.Now().UnixNano())
	post.Id = strconv.FormatInt(int64(rand.Uint64()), 10)
	// Generating time stamp
	post.PostedTimestamp = time.Now()

	// Inserting post into the database
	_, err := h.postCollection.InsertOne(context.Background(), post)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	// Post successfully created
	w.Write([]byte("successfully created post"))
}

func (h *PostHandler) getPost(w http.ResponseWriter, r *http.Request) {
	//getting id from url
	id := r.URL.Path[len("/posts/"):]
	fmt.Println(id)

	post := &data.OutPost{}
	// getting post from database
	err := h.postCollection.FindOne(context.Background(), bson.D{{"_id", id}}).Decode(post)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// sending the post to user
	functions.WriteJson(w, r, post)
}
