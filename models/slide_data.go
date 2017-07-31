package models

import (
	"gopkg.in/mgo.v2"
)

type SlideData struct {
	collection *mgo.Collection
	Id         string   `json:"id" bson:"_id"`
	Comments   []Comment `json:"comments" bson:"comments"`
	PageNumber int      `json:"page_number" bson:"page_number"`
	ImageUrl   string   `json:"image_url" bson:"image_url"`
	MarkDown   string   `json:"mark_down" bson:"mark_down"`
	Caption    string   `json:"caption" bson:"caption"`
}
