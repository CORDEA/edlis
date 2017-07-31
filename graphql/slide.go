package graphql

import (
	"github.com/graphql-go/graphql"
	"edlis/models"
)

var Slide = graphql.NewObject(graphql.ObjectConfig{
	Name: "Slide",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if slide, ok := p.Source.(models.Slide); ok {
					return slide.Id, nil
				}
				return nil, nil
			},
		},
		"title": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if slide, ok := p.Source.(models.Slide); ok {
					return slide.Title, nil
				}
				return nil, nil
			},
		},
		"description": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if slide, ok := p.Source.(models.Slide); ok {
					return slide.Description, nil
				}
				return nil, nil
			},
		},
		"slides": &graphql.Field{
			Type: graphql.NewList(SlideData),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if slide, ok := p.Source.(models.Slide); ok {
					return slide.Slides, nil
				}
				return []interface{}{}, nil
			},
		},
		"slide_count": &graphql.Field{
			Type: graphql.Int,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if slide, ok := p.Source.(models.Slide); ok {
					return slide.SlideCount, nil
				}
				return nil, nil
			},
		},
		"is_live": &graphql.Field{
			Type: graphql.Boolean,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if slide, ok := p.Source.(models.Slide); ok {
					return slide.IsLive, nil
				}
				return nil, nil
			},
		},
		"lang": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if slide, ok := p.Source.(models.Slide); ok {
					return slide.Lang, nil
				}
				return nil, nil
			},
		},
		"created_at": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if slide, ok := p.Source.(models.Slide); ok {
					return slide.CreatedAt, nil
				}
				return nil, nil
			},
		},
		"updated_at": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if slide, ok := p.Source.(models.Slide); ok {
					return slide.UpdatedAt, nil
				}
				return nil, nil
			},
		},
	},
})
