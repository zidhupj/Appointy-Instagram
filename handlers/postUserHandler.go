package handlers

import (
	"Appointy-Instagram/data"
	"Appointy-Instagram/functions"
	"context"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

			// getting query parameters of pagmentation
			limit, offset, err1 := functions.GetLimitAndOffset(w, r)
			if err1 != nil {
				http.Error(w, err1.Error(), http.StatusBadRequest)
			}

			// finding a set of posts of the user from databse
			postCursor, err := h.postCollection.Find(context.Background(), bson.D{{"userId", userId}}, &options.FindOptions{
				// Limiting the number of data retrieved from the database at a time
				Limit: &limit,
				Skip:  &offset,
			})
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
