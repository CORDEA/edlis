package models

import (
	"gopkg.in/mgo.v2"
)

type Slide struct {
	collection  *mgo.Collection
	Id          string      `json:"id" bson:"_id"`
	Title       string      `json:"title" bson:"title"`
	Description string      `json:"description" bson:"description"`
	Slides      []SlideData `json:"slides" bson:"slides"`
	SlideCount  int         `json:"slide_count" bson:"slide_count"`
	IsLive      bool        `json:"is_live" bson:"is_live"`
	Lang        string      `json:"lang" bson:"lang"`
	CreatedAt   string      `json:"created_at" bson:"created_at"`
	UpdatedAt   string      `json:"updated_at" bson:"updated_at"`
}
