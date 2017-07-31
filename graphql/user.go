package graphql

import (
	"github.com/graphql-go/graphql"
	"edlis/models"
)

var User = graphql.NewObject(graphql.ObjectConfig{
	Name: "User",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if user, ok := p.Source.(models.User); ok {
					return user.Id, nil
				}
				return nil, nil
			},
		},
		"slides": &graphql.Field{
			Type: graphql.NewList(Slide),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if user, ok := p.Source.(models.User); ok {
					return user.Slides, nil
				}
				return []interface{}{}, nil
			},
		},
		"comments": &graphql.Field{
			Type: graphql.NewList(Comment),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if user, ok := p.Source.(models.User); ok {
					return user.Comments, nil
				}
				return []interface{}{}, nil
			},
		},
		"name": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if user, ok := p.Source.(models.User); ok {
					return user.Name, nil
				}
				return nil, nil
			},
		},
		"bio": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if user, ok := p.Source.(models.User); ok {
					return user.Bio, nil
				}
				return nil, nil
			},
		},
		"avatar_url": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if user, ok := p.Source.(models.User); ok {
					return user.AvatarUrl, nil
				}
				return nil, nil
			},
		},
		"url": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if user, ok := p.Source.(models.User); ok {
					return user.Url, nil
				}
				return nil, nil
			},
		},
		"github_id": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if user, ok := p.Source.(models.User); ok {
					return user.GitHubId, nil
				}
				return nil, nil
			},
		},
		"twitter_id": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if user, ok := p.Source.(models.User); ok {
					return user.TwitterId, nil
				}
				return nil, nil
			},
		},
		"created_at": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if user, ok := p.Source.(models.User); ok {
					return user.CreatedAt, nil
				}
				return nil, nil
			},
		},
		"updated_at": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if user, ok := p.Source.(models.User); ok {
					return user.UpdatedAt, nil
				}
				return nil, nil
			},
		},
	},
})
