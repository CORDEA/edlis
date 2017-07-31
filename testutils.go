package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"edlis/client"
	"regexp"
)

func Path(query string, isMutation bool) string {
	reg := regexp.MustCompile(`[\n\s]`)
	m := ""
	if isMutation {
		m = "mutation+_"
	}

	return fmt.Sprintf("/graphql?query=%s{%s}",
		m, reg.ReplaceAllString(query, ""))
}

func Parse(response *http.Response, model interface{}) {
	data, _ := ioutil.ReadAll(response.Body)
	err := json.Unmarshal(data, model)
	if err != nil {
		log.Fatalln(err)
	}
}

func MockTestServer() *httptest.Server {
	return testServer(true)
}

func TestServer() *httptest.Server {
	return testServer(false)
}

func testServer(isTesting bool) *httptest.Server {
	c := client.MongoDb{}
	c.Open()

	mux := http.NewServeMux()
	g := client.NewGraphQl(c, isTesting)
	mux.HandleFunc("/graphql", g.Handler())

	server := httptest.NewServer(mux)
	return server
}
