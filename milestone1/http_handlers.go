package milestone1

import (
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
	"strings"
  "embed"
)

const welcomeMessage = "Welcome to support. My name is Rheo. How can I help you today?"

//go:embed index.html
var indexHtml string

var upgrader = websocket.Upgrader{}

func RootHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, string(indexHtml))
}

func ChatHandler(writer http.ResponseWriter, request *http.Request) {
	conn, err := upgrader.Upgrade(writer, request, nil)
	if err != nil {
		fmt.Println("ws connection upgrade failed", err)
		http.Error(writer, "ws connection upgrade failed", 1)
		return
	}
	conn.WriteMessage(websocket.TextMessage, []byte(welcomeMessage))

	for {
		_, bMessage, err := conn.ReadMessage()
		message := string(bMessage[:])

		if err != nil {
			fmt.Println("err")
		}
		resMessage := strings.ToUpper(message)
		conn.WriteMessage(websocket.TextMessage, []byte(resMessage))
	}
	fmt.Println("ws connection upgraded succesfully")
}

//go:embed static/*
var content embed.FS
var staticHandler = http.FileServer(http.FS(content))
