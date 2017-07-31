package main

import (
	"net/http"
	"edlis/models"
	"testing"
)

func TestSlides(t *testing.T) {
	server := MockTestServer()
	defer server.Close()

	query := `
	slides {
		id,
		title,
		description,
		slides {
			id,
			comments {
				id,
				slide_id,
				user_id,
				lang,
				comment,
				type,
				posted_at
			},
			page_number,
			image_url,
			mark_down,
			caption
		},
		slide_count,
		is_live,
		lang,
		created_at,
		updated_at
	}
	`

	response, _ := http.Get(server.URL + Path(query, false))

	var obj map[string]map[string][]models.Slide
	Parse(response, &obj)
	slides := obj["data"]["slides"]

	if len(slides) != 3 {
		t.Fatal(slides)
	}

	slide1 := slides[0]
	if slide1.Id != "1" ||
		slide1.Title != "title" ||
		slide1.Description != "description" ||
		len(slide1.Slides) != slide1.SlideCount ||
		slide1.IsLive ||
		slide1.Lang != "ja" ||
		slide1.CreatedAt != "2017-03-30T16:44:15+09:00" ||
		slide1.UpdatedAt != "2017-03-30T16:44:15+09:00" {
		t.Fatal(slide1)
	}
	slideData := slide1.Slides[0]
	if slideData.Id != "1" ||
		slideData.PageNumber != 1 ||
		slideData.ImageUrl != "image_url" ||
		slideData.MarkDown != "mark_down" ||
		slideData.Caption != "caption" {
		t.Fatal(slideData)
	}
}

func TestSlide(t *testing.T) {
	server := MockTestServer()
	defer server.Close()

	query := `
	slide {
		id,
		title,
		description,
		slides {
			id,
			comments {
				id,
				slide_id,
				user_id,
				lang,
				comment,
				type,
				posted_at
			},
			page_number,
			image_url,
			mark_down,
			caption
		},
		slide_count,
		is_live,
		lang,
		created_at,
		updated_at
	}
	`
	response, _ := http.Get(server.URL + Path(query, false))

	var obj map[string]map[string]models.Slide
	Parse(response, &obj)
	slide := obj["data"]["slide"]

	if slide.Id != "1" ||
		slide.Title != "title" ||
		slide.Description != "description" ||
		len(slide.Slides) != slide.SlideCount ||
		slide.IsLive ||
		slide.Lang != "ja" ||
		slide.CreatedAt != "2017-03-30T16:44:15+09:00" ||
		slide.UpdatedAt != "2017-03-30T16:44:15+09:00" {
		t.Fatal(slide)
	}
	slideData := slide.Slides[0]
	if slideData.Id != "1" ||
		slideData.PageNumber != 1 ||
		slideData.ImageUrl != "image_url" ||
		slideData.MarkDown != "mark_down" ||
		slideData.Caption != "caption" {
		t.Fatal(slideData)
	}
}

func TestComments(t *testing.T) {
	server := MockTestServer()
	defer server.Close()

	query := `
	comments {
		id,
		slide_id,
		user_id,
		lang,
		comment,
		type,
		posted_at
	}
	`

	response, _ := http.Get(server.URL + Path(query, false))

	var obj map[string]map[string][]models.Comment
	Parse(response, &obj)
	comments := obj["data"]["comments"]

	if len(comments) != 4 {
		t.Fatal(comments)
	}

	comment1 := comments[0]
	if comment1.Id != "1" ||
		comment1.SlideId != "1" ||
		comment1.UserId != "1" ||
		comment1.Lang != "ja" ||
		comment1.Comment != "comment" ||
		comment1.Type != models.STANDARD ||
		comment1.PostedAt != "2017-03-30T16:44:15+09:00" {
		t.Fatal(comment1)
	}
}

func TestUsers(t *testing.T) {
	server := MockTestServer()
	defer server.Close()

	query := `
	users {
		id,
		slides {
			id,
			title,
			description,
			slides {
				id,
				comments {
					id,
					slide_id,
					user_id,
					lang,
					comment,
					type,
					posted_at
				},
				page_number,
				image_url,
				mark_down,
				caption
			},
			slide_count,
			is_live,
			lang,
			created_at,
			updated_at
		},
		comments {
			id,
			slide_id,
			user_id,
			lang,
			comment,
			type,
			posted_at
		},
		name,
		bio,
		avatar_url,
		url,
		github_id,
		twitter_id,
		created_at,
		updated_at
	}
	`

	response, _ := http.Get(server.URL + Path(query, false))

	var obj map[string]map[string][]models.User
	Parse(response, &obj)
	users := obj["data"]["users"]

	if len(users) != 4 {
		t.Fatal(users)
	}

	user1 := users[0]
	if user1.Id != "1" ||
		user1.Name != "name" ||
		user1.Bio != "bio" ||
		user1.AvatarUrl != "avatar_url" ||
		user1.Url != "url" ||
		user1.GitHubId != "github_id" ||
		user1.TwitterId != "twitter_id" ||
		user1.CreatedAt != "2017-03-30T16:44:15+09:00" ||
		user1.UpdatedAt != "2017-03-30T16:44:15+09:00" {
		t.Fatal(user1)
	}
}

func TestUser(t *testing.T) {
	server := MockTestServer()
	defer server.Close()

	query := `
	user {
		id,
		slides {
			id,
			title,
			description,
			slides {
				id,
				comments {
					id,
					slide_id,
					user_id,
					lang,
					comment,
					type,
					posted_at
				},
				page_number,
				image_url,
				mark_down,
				caption
			},
			slide_count,
			is_live,
			lang,
			created_at,
			updated_at
		},
		comments {
			id,
			slide_id,
			user_id,
			lang,
			comment,
			type,
			posted_at
		},
		name,
		bio,
		avatar_url,
		url,
		github_id,
		twitter_id,
		created_at,
		updated_at
	}
	`

	response, _ := http.Get(server.URL + Path(query, false))

	var obj map[string]map[string]models.User
	Parse(response, &obj)
	user := obj["data"]["user"]

	if user.Id != "1" ||
		len(user.Comments) != 0 ||
		len(user.Slides) != 0 ||
		user.Name != "name" ||
		user.Bio != "bio" ||
		user.AvatarUrl != "avatar_url" ||
		user.Url != "url" ||
		user.GitHubId != "github_id" ||
		user.TwitterId != "twitter_id" ||
		user.CreatedAt != "2017-03-30T16:44:15+09:00" ||
		user.UpdatedAt != "2017-03-30T16:44:15+09:00" {
		t.Fatal(user)
	}
}
