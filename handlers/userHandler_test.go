package handlers

import (
	"Appointy-Instagram/data"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestUserHandler(t *testing.T) {
	mongoClient := data.ConnectToDb()
	userCollection := mongoClient.Database("Insta").Collection("Users")

	tt := []struct {
		name       string
		method     string
		uri        string
		input      *data.InUser
		statusCode int
	}{
		{
			name:   "create user without id",
			method: http.MethodPost,
			uri:    "/users/",
			input: &data.InUser{
				Id:       "",
				Name:     "aaa",
				Email:    "aaa@aaa.aaa",
				Password: "aaa",
			},
			statusCode: http.StatusBadRequest,
		}, {
			name:   "create invalid email",
			method: http.MethodPost,
			uri:    "/users/",
			input: &data.InUser{
				Id:       "aaa",
				Name:     "aaa",
				Email:    "aaaaaa.aaa",
				Password: "aaa",
			},
			statusCode: http.StatusBadRequest,
		}, {
			name:   "create valid user",
			method: http.MethodPost,
			uri:    "/users/",
			input: &data.InUser{
				Id:       "aaa",
				Name:     "aaa",
				Email:    "aaa@aaa.aaa",
				Password: "aaa",
			},
			statusCode: http.StatusOK,
		}, {
			name:   "create same valid user",
			method: http.MethodPost,
			uri:    "/users/",
			input: &data.InUser{
				Id:       "aaa",
				Name:     "aaa",
				Email:    "aaa@aaa.aaa",
				Password: "aaa",
			},
			statusCode: http.StatusBadRequest,
		}, {
			name:       "requesting valid user",
			method:     http.MethodGet,
			uri:        "/users/aaa",
			input:      &data.InUser{},
			statusCode: http.StatusOK,
		}, {
			name:       "requesting Invalid user",
			method:     http.MethodGet,
			uri:        "/users/bbb",
			input:      &data.InUser{},
			statusCode: http.StatusBadRequest,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			jsonbody, _ := json.Marshal(tc.input)
			request := httptest.NewRequest(tc.method, tc.uri, bytes.NewReader(jsonbody))

			responseRecorder := httptest.NewRecorder()

			uh := &UserHandler{userCollection: userCollection}

			uh.ServeHTTP(responseRecorder, request)
			if responseRecorder.Code != tc.statusCode {
				t.Errorf("Want status '%d', got '%d'", tc.statusCode, responseRecorder.Code)
			}
		})
	}
}
