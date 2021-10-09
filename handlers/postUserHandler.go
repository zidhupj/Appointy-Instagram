package handlers

import (
	"Appointy-Instagram/data"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type PostUserHandler struct {
	postCollection *mongo.Collection
}

func NewPostUserHandler(col *mongo.Collection) *PostUserHandler {
	return &PostUserHandler{
		postCollection: col,
	}
}

func (h *PostUserHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		w.Write([]byte("double works"))
		userId := r.URL.Path[len("/posts/users/"):]
		fmt.Println(userId)

		postCursor, err := h.postCollection.Find(context.Background(), bson.D{{"userId", userId}})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		} else {
			posts := &[]data.Post{}
			err := postCursor.All(context.Background(), posts)

			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			} else {
				jsonPosts, err := json.Marshal(posts)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
				} else {
					if len(jsonPosts) == 0 {
						http.Error(w, "No posts found for the user", http.StatusNotFound)
					} else {
						w.Write(jsonPosts)
					}
				}
			}
		}

	default:
		w.Write([]byte("Method not implemented"))
	}

}
