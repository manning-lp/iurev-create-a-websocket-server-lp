package main

// todo: learn what are packages
// todo: learn named imports
// todo: learn about passing functions into functions. Are they copied or passed by reference?
// todo: learn http.Error status int

import (
	"fmt"
	"net/http"
  "strings"
  "github.com/gorilla/websocket"
)

const welcomeMessage = "Welcome to support. My name is Rheo. How can I help you today?"

var upgrader = websocket.Upgrader{}

func handler(writer http.ResponseWriter, request *http.Request) {
  conn, err := upgrader.Upgrade(writer, request, nil)
  if err != nil {
    fmt.Println("ws connection upgrade failed", err)
    http.Error(writer, "ws connection upgrade failed", 1)
    return
  }
  // messageType, p, err := conn.ReadMessage()
  conn.WriteMessage(websocket.TextMessage, []byte(welcomeMessage));

  for {
    _, bMessage, err := conn.ReadMessage()
    message := string(bMessage[:])
    
    if err != nil {
      fmt.Println("err")
    }
    resMessage := strings.ToUpper(message)
    conn.WriteMessage(websocket.TextMessage, []byte(resMessage));
    fmt.Println(message)
  }
  fmt.Println("ws connection upgraded succesfully")
}

func main() {
  http.Handle("/", http.FileServer(http.Dir("./milestone1")))
	http.HandleFunc("/chat", handler)

	http.ListenAndServe(":8080", nil) // todo: learn about second arg
}
