package main

import (
	"golang.org/x/net/websocket"
	"net/http"
	"edlis/client"
)

func main() {
	c := client.MongoDb{}
	c.Open()

	g := client.NewGraphQl(c, false)
	h := client.WebSocket{}

	defer c.Session.Close()

	http.HandleFunc("/graphql", g.Handler())
	http.Handle("/chat", websocket.Handler(h.Handler()))
	http.ListenAndServe(":8080", nil)
}
