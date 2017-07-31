package db

import (
	"gopkg.in/mgo.v2"
	"log"
	"edlis/models"
	"gopkg.in/mgo.v2/bson"
)

type SlideData struct {
	Id         string `bson:"_id"`
	CommentIds []string      `bson:"comment_ids"`
	PageNumber int           `bson:"page_number"`
	ImageUrl   string        `bson:"image_url"`
	MarkDown   string        `bson:"mark_down"`
	Caption    string        `bson:"caption"`
}

func NewSlideData(pageNumber int, markDown string) SlideData {
	return SlideData{
		Id:         bson.NewObjectId().Hex(),
		CommentIds: []string{},
		PageNumber: pageNumber,
		ImageUrl:   "",
		MarkDown:   markDown,
		Caption:    "",
	}
}

func (s *SlideData) ToSlideData(c *mgo.Collection) models.SlideData {
	var comments []models.Comment
	for _, w := range s.CommentIds {
		var cs []models.Comment
		if err := c.FindId(w).All(&cs); err != nil {
			log.Fatalln(err)
		}
		comments = append(comments, cs[0])
	}

	return models.SlideData{
		Id:         s.Id,
		Comments:   comments,
		PageNumber: s.PageNumber,
		ImageUrl:   s.ImageUrl,
		MarkDown:   s.MarkDown,
		Caption:    s.Caption,
	}
}
