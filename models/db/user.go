package db

import (
	"gopkg.in/mgo.v2"
	"log"
	"edlis/models"
	"gopkg.in/mgo.v2/bson"
	"github.com/graphql-go/graphql"
)

type User struct {
	Id         string `bson:"_id"`
	SlideIds   []string      `bson:"slide_ids"`
	CommentIds []string      `bson:"comment_ids"`
	Name       string        `bson:"name"`
	Bio        string        `bson:"bio"`
	AvatarUrl  string        `bson:"avatar_url"`
	Url        string        `bson:"url"`
	GitHubId   string        `bson:"github_id"`
	TwitterId  string        `bson:"twitter_id"`
	CreatedAt  string        `bson:"created_at"`
	UpdatedAt  string        `bson:"updated_at"`
}

func ToUser(p graphql.ResolveParams) User {
	return newUser(
		p.Args["name"].(string),
		p.Args["bio"].(string),
		p.Args["avatar_url"].(string),
		p.Args["url"].(string),
		p.Args["github_id"].(string),
		p.Args["twitter_id"].(string),
	)
}

func newUser(name string, bio string, avatarUrl string,
	url string, gitHubUrl string, twitterId string) User {
	return User{
		Id:        bson.NewObjectId().Hex(),
		Name:      name,
		Bio:       bio,
		AvatarUrl: avatarUrl,
		Url:       url,
		GitHubId:  gitHubUrl,
		TwitterId: twitterId,
		CreatedAt: FormattedTime(),
		UpdatedAt: FormattedTime(),
	}
}

func (u *User) Update(p graphql.ResolveParams) {
	u.Name = p.Args["name"].(string)
	u.Bio = p.Args["bio"].(string)
	u.AvatarUrl = p.Args["avatar_url"].(string)
	u.Url = p.Args["url"].(string)
	u.GitHubId = p.Args["github_id"].(string)
	u.TwitterId = p.Args["twitter_id"].(string)
	u.UpdatedAt = FormattedTime()
}

func (u *User) ToUser(comment *mgo.Collection, slide *mgo.Collection) models.User {
	var slides []models.Slide
	var comments []models.Comment
	for _, w := range u.CommentIds {
		var cs []models.Comment
		if err := comment.FindId(w).All(&cs); err != nil {
			log.Fatalln(err)
		}
		comments = append(comments, cs[0])
	}
	for _, w := range u.SlideIds {
		var ss []models.Slide
		if err := slide.FindId(w).All(&ss); err != nil {
			log.Fatalln(err)
		}
		slides = append(slides, ss[0])
	}

	return models.User{
		Id:        u.Id,
		Slides:    slides,
		Comments:  comments,
		Name:      u.Name,
		Bio:       u.Bio,
		AvatarUrl: u.AvatarUrl,
		Url:       u.Url,
		GitHubId:  u.GitHubId,
		TwitterId: u.TwitterId,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}
