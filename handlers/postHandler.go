package handlers

import (
	"Appointy-Instagram/data"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
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
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	post := &data.Post{}
	json.Unmarshal(body, post)

	fmt.Println(post)

	rand.Seed(time.Now().UnixNano())
	post.Id = strconv.FormatInt(int64(rand.Uint64()), 10)
	post.PostedTimestamp = time.Now()

	err = nil
	_, err = h.postCollection.InsertOne(context.Background(), post)
	if err != nil {
		w.Write([]byte(err.Error()))
	} else {
		w.Write([]byte("successfully created post"))
	}

}

func (h *PostHandler) getPost(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/posts/"):]
	fmt.Println(id)

	post := &data.Post{}

	err := h.postCollection.FindOne(context.Background(), bson.D{{"_id", id}}).Decode(post)
	if err != nil {
		w.Write([]byte("Post does not exist"))
	} else {
		jsonPost, err := json.Marshal(post)
		if err != nil {
			w.Write([]byte("Unable to marshal post"))
		} else {
			w.Write(jsonPost)
		}
	}
}
