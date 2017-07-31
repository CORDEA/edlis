package main

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"edlis/models"
	"fmt"
)

func TestSlideWithDb(t *testing.T) {
	server := TestServer()
	defer server.Close()

	slide, ok := AddSlide(server)
	if !ok {
		t.Fatal(slide)
	}

	slide, ok = GetSlide(server, slide)
	if !ok {
		t.Fatal(slide)
	}

	slide, ok = UpdateSlide(server, slide.Id)
	if !ok {
		t.Fatal(slide)
	}

	if ok = DeleteSlide(server, slide.Id); !ok {
		t.Fatal(slide)
	}
}

func GetSlide(server *httptest.Server, expected models.Slide) (models.Slide, bool) {
	query := fmt.Sprintf(`
	slide(id: "%s") {
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
	`, expected.Id)
	response, _ := http.Get(server.URL + Path(query, false))

	var obj map[string]map[string]models.Slide
	Parse(response, &obj)
	slide := obj["data"]["slide"]

	if slide.Title != expected.Title ||
		slide.Description != expected.Description ||
		len(slide.Slides) != slide.SlideCount {
		return slide, false
	}
	slideData := slide.Slides[0]
	if slideData.PageNumber != expected.Slides[0].PageNumber ||
		slideData.MarkDown != expected.Slides[0].MarkDown {
		return slide, false
	}
	return slide, true
}

func AddSlide(server *httptest.Server) (models.Slide, bool) {
	query := `
	addSlide(title: "title", description: "description", slides: [
		{ page_number: 1, mark_down: "mark_down" }
	]) {
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
	if slide.Title != "title" ||
		slide.Description != "description" {
		return slide, false
	}
	return slide, true
}

func UpdateSlide(server *httptest.Server, id string) (models.Slide, bool) {
	query := fmt.Sprintf(`
	updateSlide(id: "%s", title: "title2", description: "description2", slides: []) {
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
	`, id)

	response, _ := http.Get(server.URL + Path(query, true))
	var obj map[string]map[string]models.Slide
	Parse(response, &obj)

	slide := obj["data"]["updateSlide"]
	if slide.Title != "title2" ||
		slide.Description != "description2" {
		return slide, false
	}
	return slide, true
}

func DeleteSlide(server *httptest.Server, id string) bool {
	query := fmt.Sprintf(`
	deleteSlide(id: "%s") {
		id
	}
	`, id)

	response, _ := http.Get(server.URL + Path(query, true))
	var obj map[string]map[string]models.Id
	Parse(response, &obj)

	deletedId := obj["data"]["deleteSlide"]
	if deletedId.Id != id {
		return false
	}
	return true
}
