package client

import (
	"github.com/graphql-go/graphql"
	"gopkg.in/mgo.v2/bson"
	"log"
	objects "edlis/graphql"
	"edlis/models"
	"edlis/models/db"
)

func Mutations(isTesting bool, client MongoDb) *graphql.Object {

	var updateSlideType = graphql.NewInputObject(graphql.InputObjectConfig{
		Name: "Slide",
		Fields: graphql.InputObjectConfigFieldMap{
			"page_number": &graphql.InputObjectFieldConfig{
				Type: graphql.Int,
			},
			"mark_down": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
		},
	})

	return graphql.NewObject(graphql.ObjectConfig{
		Name: "RootMutation",
		Fields: graphql.Fields{
			"addSlide": &graphql.Field{
				Type:        objects.Slide,
				Description: "",
				Args: graphql.FieldConfigArgument{
					"title": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"description": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"slides": &graphql.ArgumentConfig{
						Type: graphql.NewList(updateSlideType),
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if isTesting {
						slide := models.Slide{}
						ParseTestJson("addSlide", &slide)
						return slide, nil
					}

					slide := db.ToSlide(p)
					if err := client.Slide().Insert(slide); err != nil {
						log.Fatalln(err)
					}

					s := slide.ToSlide(client.SlideData())

					return s, nil
				},
			},
			"updateSlide": &graphql.Field{
				Type:        objects.Slide,
				Description: "",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"title": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"description": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"slides": &graphql.ArgumentConfig{
						Type: graphql.NewList(updateSlideType),
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if isTesting {
						slide := models.Slide{}
						ParseTestJson("updateSlide", &slide)
						return slide, nil
					}
					id := p.Args["id"].(string)
					c := client.Slide()

					var slide db.Slide
					if err := c.FindId(id).One(&slide); err != nil {
						log.Fatalln(err)
					}

					slide.Update(p)
					c.UpdateId(id, slide)

					return slide.ToSlide(client.SlideData()), nil
				},
			},
			"deleteSlide": &graphql.Field{
				Type:        objects.Id,
				Description: "",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if isTesting {
						id := models.Id{}
						ParseTestJson("deleteSlide", &id)
						return id, nil
					}
					id := p.Args["id"].(string)
					client.Slide().RemoveId(id)
					return models.NewId(id), nil
				},
			},
			"addComment": &graphql.Field{
				Type:        objects.Comment,
				Description: "",
				Args: graphql.FieldConfigArgument{
					"slide_id": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"comment": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"type": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if isTesting {
						comment := models.Comment{}
						ParseTestJson("addComment", &comment)
						return comment, nil
					}
					comment := models.Comment{
						Id:      bson.NewObjectId().Hex(),
						SlideId: p.Args["slide_id"].(string),
						Comment: p.Args["comment"].(string),
						// TODO: add userId
						UserId: "1",
						// TODO:
						Lang: "ja",
						// TODO:
						Type:     p.Args["type"].(string),
						PostedAt: db.FormattedTime(),
					}
					c := client.Comment()
					if err := c.Insert(&comment); err != nil {
						log.Fatalln(err)
					}

					s := client.Slide()
					slide := db.Slide{}
					if err := s.FindId(comment.SlideId).One(&slide); err != nil {
						for _, data := range slide.Slides {
							if data.Id == comment.SlideId {
								data.CommentIds = append(data.CommentIds, comment.Id)
								break
							}
						}
						s.UpdateId(slide.Id, slide)
					}

					u := client.User()
					user := db.User{}
					if err := u.FindId(comment.UserId).One(&user); err != nil {
						user.CommentIds = append(user.CommentIds, comment.Id)
						u.UpdateId(user.Id, user)
					}

					return comment, nil
				},
			},
			"createUser": &graphql.Field{
				Type:        objects.User,
				Description: "",
				Args: graphql.FieldConfigArgument{
					"name": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"bio": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"avatar_url": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"url": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"github_id": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"twitter_id": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if isTesting {
						user := models.User{}
						ParseTestJson("createUser", &user)
						return user, nil
					}

					user := db.ToUser(p)
					client.User().Insert(user)
					return user.ToUser(client.Comment(), client.Slide()), nil
				},
			},
			"updateUser": &graphql.Field{
				Type:        objects.User,
				Description: "",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"name": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"bio": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"avatar_url": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"url": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"github_id": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"twitter_id": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if isTesting {
						user := models.User{}
						ParseTestJson("updateUser", &user)
						return user, nil
					}
					id := p.Args["id"].(string)
					c := client.User()
					var user db.User
					if err := c.FindId(id).One(&user); err != nil {
						log.Fatalln(err)
					}

					user.Update(p)
					c.UpdateId(id, user)

					return user.ToUser(client.Comment(), client.Slide()), nil
				},
			},
			"deleteUser": &graphql.Field{
				Type:        objects.Id,
				Description: "",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if isTesting {
						id := models.Id{}
						ParseTestJson("deleteUser", &id)
						return id, nil
					}
					id := p.Args["id"].(string)
					client.User().RemoveId(id)

					return models.NewId(id), nil
				},
			},
		},
	})
}
