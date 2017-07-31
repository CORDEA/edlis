package models

import (
	"gopkg.in/mgo.v2"
)

type User struct {
	commentCollection *mgo.Collection
	slideCollection   *mgo.Collection
	Id                string   `json:"id" bson:"_id"`
	Slides            []Slide   `json:"slides" bson:"slides"`
	Comments          []Comment `json:"comments" bson:"comments"`
	Name              string `json:"name" bson:"name"`
	Bio               string `json:"bio" bson:"bio"`
	AvatarUrl         string `json:"avatar_url" bson:"avatar_url"`
	Url               string `json:"url" bson:"url"`
	GitHubId          string `json:"github_id" bson:"github_id"`
	TwitterId         string `json:"twitter_id" bson:"twitter_id"`
	CreatedAt         string `json:"created_at" bson:"created_at"`
	UpdatedAt         string `json:"updated_at" bson:"updated_at"`
}
