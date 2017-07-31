package graphql

import (
	"github.com/graphql-go/graphql"
	"edlis/models"
)

var Comment = graphql.NewObject(graphql.ObjectConfig{
	Name: "Comment",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if comment, ok := p.Source.(models.Comment); ok {
					return comment.Id, nil
				}
				return nil, nil
			},
		},
		"slide_id": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if comment, ok := p.Source.(models.Comment); ok {
					return comment.SlideId, nil
				}
				return nil, nil
			},
		},
		"user_id": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if comment, ok := p.Source.(models.Comment); ok {
					return comment.UserId, nil
				}
				return nil, nil
			},
		},
		"lang": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if comment, ok := p.Source.(models.Comment); ok {
					return comment.Lang, nil
				}
				return nil, nil
			},
		},
		"comment": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if comment, ok := p.Source.(models.Comment); ok {
					return comment.Comment, nil
				}
				return nil, nil
			},
		},
		"type": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if comment, ok := p.Source.(models.Comment); ok {
					return comment.Type, nil
				}
				return nil, nil
			},
		},
		"posted_at": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if comment, ok := p.Source.(models.Comment); ok {
					return comment.PostedAt, nil
				}
				return nil, nil
			},
		},
	},
})
