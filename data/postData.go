package data

import (
	"time"
)

type Post struct {
	UserId          string    `json:"userId" bson:"userId"`
	Id              string    `json:"id" bson:"_id"`
	Caption         string    `json:"caption" bson:"caption"`
	ImgUrl          string    `json:"imgUrl" bson:"imgUrl"`
	PostedTimestamp time.Time `json:"postedTimestamp" bson:"postedTimestamp"`
}
