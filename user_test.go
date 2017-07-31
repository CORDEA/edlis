package main

import (
	"testing"
	"net/http/httptest"
	"edlis/models"
	"net/http"
	"fmt"
)

func TestUserWithDb(t *testing.T) {
	server := TestServer()
	defer server.Close()

	user, ok := AddUser(server)
	if !ok {
		t.Fatal(user)
	}

	user, ok = GetUser(server, user)
	if !ok {
		t.Fatal(user)
	}

	user, ok = UpdateUser(server, user.Id)
	if !ok {
		t.Fatal(user)
	}

	if ok = DeleteUser(server, user.Id); !ok {
		t.Fatal(user)
	}
}

func GetUser(server *httptest.Server, expected models.User) (models.User, bool) {
	query := fmt.Sprintf(`
	user(id: "%s") {
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
	`, expected.Id)
	response, _ := http.Get(server.URL + Path(query, false))

	var obj map[string]map[string]models.User
	Parse(response, &obj)
	user := obj["data"]["user"]

	if user.Id != expected.Id ||
		len(user.Comments) != len(expected.Comments) ||
		len(user.Slides) != len(expected.Slides) ||
		user.Name != expected.Name ||
		user.Bio != expected.Bio ||
		user.AvatarUrl != expected.AvatarUrl ||
		user.Url != expected.Url ||
		user.GitHubId != expected.GitHubId ||
		user.TwitterId != expected.TwitterId {
		return user, false
	}
	return user, true
}

func AddUser(server *httptest.Server) (models.User, bool) {
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
	if user.Name != "name" ||
		user.Bio != "bio" ||
		user.AvatarUrl != "avatar_url" ||
		user.Url != "url" ||
		len(user.Comments) != 0 ||
		len(user.Slides) != 0 ||
		user.GitHubId != "github_id" ||
		user.TwitterId != "twitter_id" {
		return user, false
	}
	return user, true
}

func UpdateUser(server *httptest.Server, id string) (models.User, bool) {
	query := fmt.Sprintf(`
	updateUser(id: "%s", name: "name2", bio: "bio2", avatar_url: "avatar_url2", url: "url2", github_id: "github_id2", twitter_id: "twitter_id2") {
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
	`, id)

	response, _ := http.Get(server.URL + Path(query, true))
	var obj map[string]map[string]models.User
	Parse(response, &obj)

	user := obj["data"]["updateUser"]
	if user.Id != id ||
		len(user.Slides) != 0 ||
		len(user.Comments) != 0 ||
		user.Name != "name2" ||
		user.Bio != "bio2" ||
		user.AvatarUrl != "avatar_url2" ||
		user.Url != "url2" ||
		user.GitHubId != "github_id2" ||
		user.TwitterId != "twitter_id2" {
		return user, false
	}
	return user, true
}

func DeleteUser(server *httptest.Server, id string) bool {
	query := fmt.Sprintf(`
	deleteUser(id: "%s") {
		id
	}
	`, id)

	response, _ := http.Get(server.URL + Path(query, true))
	var obj map[string]map[string]models.Id
	Parse(response, &obj)

	deletedId := obj["data"]["deleteUser"]
	if deletedId.Id != id {
		return false
	}
	return true
}
