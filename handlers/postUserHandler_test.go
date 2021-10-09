package handlers

import (
	"Appointy-Instagram/data"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPostUserHandler(t *testing.T) {
	mongoClient := data.ConnectToDb()
	postCollection := mongoClient.Database("Insta").Collection("Posts")

	tt := []struct {
		name       string
		method     string
		uri        string
		statusCode int
	}{
		{
			name:       "valid request",
			method:     http.MethodGet,
			uri:        "/posts/users/bobby?limit=2&offset=1",
			statusCode: http.StatusOK,
		}, {
			name:       "without limit",
			method:     http.MethodGet,
			uri:        "/posts/users/bobby?offset=1",
			statusCode: http.StatusBadRequest,
		}, {
			name:       "limit greater than 20",
			method:     http.MethodGet,
			uri:        "/posts/users/bobby?limit=25&offset=1",
			statusCode: http.StatusBadRequest,
		}, {
			name:       "posts exhausted",
			method:     http.MethodGet,
			uri:        "/posts/users/bobby?limit=2&offset=10",
			statusCode: http.StatusNotFound,
		}, {
			name:       "user does not exist or user with no post",
			method:     http.MethodGet,
			uri:        "/posts/users/tony?limit=2&offset=10",
			statusCode: http.StatusNotFound,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			request := httptest.NewRequest(tc.method, tc.uri, nil)

			responseRecorder := httptest.NewRecorder()

			uh := &PostUserHandler{postCollection: postCollection}

			uh.ServeHTTP(responseRecorder, request)
			if responseRecorder.Code != tc.statusCode {
				t.Errorf("Want status '%d', got '%d'", tc.statusCode, responseRecorder.Code)
			}
		})
	}
}
