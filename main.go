package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"golang.org/x/net/websocket"
)

func main() {
	http.HandleFunc("/", mainHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.Handle("/ws", websocket.Handler(handleWebSocket))

	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}

func mainHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("template/index.html")
	if err != nil {
		log.Fatal(err)
	}
	if err := t.Execute(w, nil); err != nil {
		log.Fatal(err)
	}
}

func handleWebSocket(ws *websocket.Conn) {
	defer ws.Close()

	if err := websocket.Message.Send(ws, "Gopherくんのこんにちは"); err != nil {
		log.Fatal(err)
	}

	for {
		var msg string
		if err := websocket.Message.Receive(ws, &msg); err != nil {
			log.Fatal(err)
		}

		if err := websocket.Message.Send(ws, fmt.Sprintf("メッセージありがとう！: %s", msg)); err != nil {
			log.Fatal(err)
		}
	}
}
