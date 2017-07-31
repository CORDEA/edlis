package main

import (
	"net/http"
	"edlis/models"
	"testing"
)

func TestMockAddSlide(t *testing.T) {
	server := MockTestServer()
	defer server.Close()

	query := `
	addSlide(title: "title", description: "description", slides: []) {
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

	response, _ := http.Get(server.URL + Path(query, true))
	var obj map[string]map[string]models.Slide
	Parse(response, &obj)

	slide := obj["data"]["addSlide"]
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

func TestMockUpdateSlide(t *testing.T) {
	server := MockTestServer()
	defer server.Close()

	query := `
	updateSlide(id: "id", title: "title", description: "description", slides: []) {
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

	response, _ := http.Get(server.URL + Path(query, true))
	var obj map[string]map[string]models.Slide
	Parse(response, &obj)

	slide := obj["data"]["updateSlide"]
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

func TestMockDeleteSlide(t *testing.T) {
	server := MockTestServer()
	defer server.Close()

	query := `
	deleteSlide(id: "id") {
		id
	}
	`

	response, _ := http.Get(server.URL + Path(query, true))
	var obj map[string]map[string]models.Id
	Parse(response, &obj)

	id := obj["data"]["deleteSlide"]
	if id.Id != "1" {
		t.Fatal(id)
	}
}

func TestMockAddComment(t *testing.T) {
	server := MockTestServer()
	defer server.Close()

	query := `
	addComment(slide_id: "slide_id", comment: "comment", type: "type") {
		id,
		slide_id,
		user_id,
		lang,
		comment,
		type,
		posted_at
	}
	`

	response, _ := http.Get(server.URL + Path(query, true))
	var obj map[string]map[string]models.Comment
	Parse(response, &obj)

	comment := obj["data"]["addComment"]
	if comment.Id != "1" ||
		comment.SlideId != "1" ||
		comment.UserId != "1" ||
		comment.Lang != "ja" ||
		comment.Comment != "comment" ||
		comment.Type != models.STANDARD ||
		comment.PostedAt != "2017-03-30T16:44:15+09:00" {
		t.Fatal(comment)
	}
}

func TestMockCreateUser(t *testing.T) {
	server := MockTestServer()
	defer server.Close()

	query := `
	createUser(name: "name", bio: "bio", avatar_url: "avatar_url", url: "url", github_id: "github_id", twitter_id: "twitter_id") {
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

	response, _ := http.Get(server.URL + Path(query, true))
	var obj map[string]map[string]models.User
	Parse(response, &obj)

	user := obj["data"]["createUser"]
	if user.Id != "1" ||
		user.Name != "name" ||
		user.Bio != "bio" ||
		user.AvatarUrl != "avatar_url" ||
		user.Url != "url" ||
		len(user.Comments) != 1 ||
		len(user.Slides) != 0 ||
		user.GitHubId != "github_id" ||
		user.TwitterId != "twitter_id" ||
		user.CreatedAt != "2017-03-30T16:44:15+09:00" ||
		user.UpdatedAt != "2017-03-30T16:44:15+09:00" {
		t.Fatal(user)
	}
}

func TestMockUpdateUser(t *testing.T) {
	server := MockTestServer()
	defer server.Close()

	query := `
	updateUser(id: "id", name: "name", bio: "bio", avatar_url: "avatar_url", url: "url", github_id: "github_id", twitter_id: "twitter_id") {
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

	response, _ := http.Get(server.URL + Path(query, true))
	var obj map[string]map[string]models.User
	Parse(response, &obj)

	user := obj["data"]["updateUser"]
	if user.Id != "1" ||
		len(user.Slides) != 0 ||
		len(user.Comments) != 0 ||
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

func TestMockDeleteUser(t *testing.T) {
	server := MockTestServer()
	defer server.Close()

	query := `
	deleteUser {
		id
	}
	`

	response, _ := http.Get(server.URL + Path(query, true))
	var obj map[string]map[string]models.Id
	Parse(response, &obj)
	id := obj["data"]["deleteUser"]
	if id.Id != "1" {
		t.Fatal(id)
	}
}
