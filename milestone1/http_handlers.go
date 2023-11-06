package milestone1

import (
	"fmt"
  "os"
  "path/filepath"
	"net/http"
  "strings"
  "github.com/gorilla/websocket"
)

const welcomeMessage = "Welcome to support. My name is Rheo. How can I help you today?"

var indexHtmlPath, _ = filepath.Abs("milestone1/index.html")
var indexHtml, _ = os.ReadFile(indexHtmlPath)

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

var staticHandler = http.FileServer(http.Dir("./milestone1/static"))
