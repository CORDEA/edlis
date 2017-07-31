package db

import (
	"gopkg.in/mgo.v2"
	"edlis/models"
	"gopkg.in/mgo.v2/bson"
	"github.com/graphql-go/graphql"
)

type Slide struct {
	Id          string `bson:"_id"`
	Title       string        `bson:"title"`
	Description string        `bson:"description"`
	Slides      []SlideData   `bson:"slides"`
	SlideCount  int           `bson:"slide_count"`
	IsLive      bool          `bson:"is_live"`
	Lang        string        `bson:"lang"`
	CreatedAt   string        `bson:"created_at"`
	UpdatedAt   string        `bson:"updated_at"`
}

func ToSlide(p graphql.ResolveParams) Slide {
	slides := []SlideData{}
	for _, w := range p.Args["slides"].([]interface{}) {
		s := w.(map[string]interface{})
		slide := NewSlideData(
			s["page_number"].(int), s["mark_down"].(string))
		slides = append(slides, slide)
	}
	return newSlide(p.Args["title"].(string),
		p.Args["description"].(string), slides)
}

func newSlide(title string, description string, slides []SlideData) Slide {
	return Slide{
		Id:          bson.NewObjectId().Hex(),
		Title:       title,
		Description: description,
		Slides:      slides,
		SlideCount:  len(slides),
		IsLive:      false,
		Lang:        "",
		CreatedAt:   FormattedTime(),
		UpdatedAt:   FormattedTime(),
	}
}

func (s *Slide) Update(p graphql.ResolveParams) {
	s.Title = p.Args["title"].(string)
	s.Description = p.Args["description"].(string)
	s.UpdatedAt = FormattedTime()

	updateSlides := map[string]map[string]interface{}{}
	for _, w := range p.Args["slides"].([]interface{}) {
		s := w.(map[string]interface{})
		if v, ok := s["id"]; ok {
			updateSlides[v.(string)] = s
		} else {
			updateSlides[bson.NewObjectId().Hex()] = s
		}
	}

	slides := s.Slides
	for _, data := range slides {
		if v, ok := updateSlides[s.Id]; ok {
			data.PageNumber = v["page_number"].(int)
			data.MarkDown = v["mark_down"].(string)
			delete(updateSlides, s.Id)
		}
	}

	if len(updateSlides) > 0 {
		for _, w := range updateSlides {
			slide := SlideData{
				Id:         bson.NewObjectId().Hex(),
				PageNumber: w["page_number"].(int),
				MarkDown:   w["mark_down"].(string),
			}
			slides = append(slides, slide)
		}
	}

	s.Slides = slides
}

func (s *Slide) ToSlide(c *mgo.Collection) models.Slide {
	var slides []models.SlideData
	for _, w := range s.Slides {
		slides = append(slides, w.ToSlideData(c))
	}

	return models.Slide{
		Id:          s.Id,
		Title:       s.Title,
		Description: s.Description,
		Slides:      slides,
		SlideCount:  s.SlideCount,
		IsLive:      s.IsLive,
		Lang:        s.Lang,
		CreatedAt:   s.CreatedAt,
		UpdatedAt:   s.UpdatedAt,
	}
}
