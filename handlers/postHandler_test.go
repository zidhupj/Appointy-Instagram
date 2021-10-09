package handlers

import (
	"Appointy-Instagram/data"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPostHandler(t *testing.T) {
	mongoClient := data.ConnectToDb()
	postCollection := mongoClient.Database("Insta").Collection("Posts")

	tt := []struct {
		name       string
		method     string
		uri        string
		input      *data.InPost
		statusCode int
	}{
		{
			name:   "create valid post",
			method: http.MethodPost,
			uri:    "/posts/",
			input: &data.InPost{
				UserId:  "hello",
				Caption: "hiking",
				ImgUrl:  "https://external-content.duckduckgo.com/iu/?u=https%3A%2F%2Ftse1.mm.bing.net%2Fth%3Fid%3DOIP.zatSKe_nXrKH2uEtUT1ElgHaE8%26pid%3DApi&f=1",
			},
			statusCode: http.StatusOK,
		}, {
			name:   "create invalid post",
			method: http.MethodPost,
			uri:    "/posts/",
			input: &data.InPost{
				UserId:  "hello",
				Caption: "hiking",
				ImgUrl:  "",
			},
			statusCode: http.StatusBadRequest,
		}, {
			name:       "get valid post",
			method:     http.MethodGet,
			uri:        "/posts/8588481130347139375",
			input:      &data.InPost{},
			statusCode: http.StatusOK,
		}, {
			name:       "get invalid post",
			method:     http.MethodGet,
			uri:        "/posts/8588481130",
			input:      &data.InPost{},
			statusCode: http.StatusBadRequest,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			jsonbody, _ := json.Marshal(tc.input)
			request := httptest.NewRequest(tc.method, tc.uri, bytes.NewReader(jsonbody))

			responseRecorder := httptest.NewRecorder()

			uh := &PostHandler{postCollection: postCollection}

			uh.ServeHTTP(responseRecorder, request)
			if responseRecorder.Code != tc.statusCode {
				t.Errorf("Want status '%d', got '%d'", tc.statusCode, responseRecorder.Code)
			}
		})
	}
}
