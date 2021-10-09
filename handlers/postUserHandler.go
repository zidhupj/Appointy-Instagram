package handlers

import (
	"Appointy-Instagram/data"
	"Appointy-Instagram/functions"
	"context"
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
		{
			// getting id from url
			userId := r.URL.Path[len("/posts/users/"):]

			// finding a set of posts of the user from databse
			postCursor, err := h.postCollection.Find(context.Background(), bson.D{{"userId", userId}})
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			posts := &[]data.OutPost{}
			err = postCursor.All(context.Background(), posts)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			functions.WriteJson(w, r, posts)
		}

	default:
		{
			w.Write([]byte("Method not implemented"))
		}
	}

}
