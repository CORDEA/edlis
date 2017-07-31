package graphql

import (
	"github.com/graphql-go/graphql"
	"edlis/models"
)

var SlideData = graphql.NewObject(graphql.ObjectConfig{
	Name: "SlideData",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if slide, ok := p.Source.(models.SlideData); ok {
					return slide.Id, nil
				}
				return nil, nil
			},
		},
		"comments": &graphql.Field{
			Type: graphql.NewList(Comment),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if slide, ok := p.Source.(models.SlideData); ok {
					return slide.Comments, nil
				}
				return []interface{}{}, nil
			},
		},
		"page_number": &graphql.Field{
			Type: graphql.Int,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if slide, ok := p.Source.(models.SlideData); ok {
					return slide.PageNumber, nil
				}
				return nil, nil
			},
		},
		"image_url": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if slide, ok := p.Source.(models.SlideData); ok {
					return slide.ImageUrl, nil
				}
				return nil, nil
			},
		},
		"mark_down": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if slide, ok := p.Source.(models.SlideData); ok {
					return slide.MarkDown, nil
				}
				return nil, nil
			},
		},
		"caption": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if slide, ok := p.Source.(models.SlideData); ok {
					return slide.Caption, nil
				}
				return nil, nil
			},
		},
	},
})
