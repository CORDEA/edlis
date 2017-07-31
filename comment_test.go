package main

import (
	"testing"
	"net/http/httptest"
	"edlis/models"
	"net/http"
)

func TestCommentWithDb(t *testing.T) {
	server := TestServer()
	defer server.Close()

	if comment, ok := AddComment(server); !ok {
		t.Fatal(comment)
	}
}

func AddComment(server *httptest.Server) (models.Comment, bool) {
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
	if comment.SlideId != "slide_id" ||
		comment.UserId != "1" ||
		comment.Comment != "comment" {
		return comment, false
	}
	return comment, true
}
