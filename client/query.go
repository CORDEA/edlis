package client

import (
	"github.com/graphql-go/graphql"
	"log"
	objects "edlis/graphql"
	"edlis/models"
	"edlis/models/db"
)

func Queries(isTesting bool, client MongoDb) *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "RootQuery",
		Fields: graphql.Fields{
			"slides": &graphql.Field{
				Type: graphql.NewList(objects.Slide),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if isTesting {
						slides := []models.Slide{}
						ParseTestJson("slides", &slides)
						return slides, nil
					}
					c := client.Slide()
					var dbSlides []db.Slide
					if err := c.FindId(nil).All(&dbSlides); err != nil {
						log.Fatalln(err)
					}
					var slides []models.Slide
					for _, s := range dbSlides {
						slides = append(slides, s.ToSlide(client.Comment()))
					}
					return slides, nil
				},
			},
			"slide": &graphql.Field{
				Type: objects.Slide,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if isTesting {
						slide := models.Slide{}
						ParseTestJson("slide", &slide)
						return slide, nil
					}
					c := client.Slide()
					var slide db.Slide
					if err := c.FindId(p.Args["id"].(string)).One(&slide); err != nil {
						log.Fatalln(err)
					}
					return slide.ToSlide(client.Comment()), nil
				},
			},
			"comments": &graphql.Field{
				Type: graphql.NewList(objects.Comment),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if isTesting {
						comments := []models.Comment{}
						ParseTestJson("comments", &comments)
						return comments, nil
					}
					c := client.Comment()
					var comments []models.Comment
					if err := c.FindId(nil).All(&comments); err != nil {
						log.Fatalln(err)
					}
					return comments, nil
				},
			},
			"users": &graphql.Field{
				Type: graphql.NewList(objects.User),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if isTesting {
						users := []models.User{}
						ParseTestJson("users", &users)
						return users, nil
					}
					c := client.User()
					var dbUsers []db.User
					if err := c.FindId(nil).All(&dbUsers); err != nil {
						log.Fatalln(err)
					}
					var users []models.User
					for _, w := range dbUsers {
						users = append(users, w.ToUser(
							client.Comment(),
							client.Slide(),
						))
					}
					return users, nil
				},
			},
			"user": &graphql.Field{
				Type: objects.User,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if isTesting {
						user := models.User{}
						ParseTestJson("user", &user)
						return user, nil
					}
					c := client.User()
					var user db.User
					if err := c.FindId(p.Args["id"].(string)).One(&user); err != nil {
						log.Fatalln(err)
					}
					return user.ToUser(
						client.Comment(),
						client.Slide(),
					), nil
				},
			},
		},
	})
}
