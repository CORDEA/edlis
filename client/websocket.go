package client

import (
	"golang.org/x/net/websocket"
	"log"
	"edlis/models"
	"edlis/models/request"
)

type Conn struct {
	conn    *websocket.Conn
	request request.Join
}

type WebSocket struct {
	conn map[string]Conn
}

func (ws *WebSocket) send(w *websocket.Conn, msg request.Join) {
	var comment models.Comment
	for {
		if err := websocket.JSON.Receive(w, &comment); err != nil {
			delete(ws.conn, msg.Id)
			break
		}
		for _, w := range ws.conn {
			if w.request.SlideId == msg.SlideId {
				if err := websocket.JSON.Send(w.conn, comment); err != nil {
					delete(ws.conn, msg.Id)
					break
				}
			}
		}
	}
}

func (ws *WebSocket) Handler() func(w *websocket.Conn) {
	return func(w *websocket.Conn) {
		var msg request.Join
		if err := websocket.JSON.Receive(w, &msg); err != nil {
			log.Fatalln(err)
		}

		ws.conn[msg.Id] = Conn{conn: w, request: msg}

		ws.send(w, msg)
	}
}
