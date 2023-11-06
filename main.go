package main

// todo: learn what are packages
// todo: learn named imports

import (
	"fmt"
	"net/http"
)


func handler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Hello World, %s", request.URL.Path[:1])
}

func main() {
  http.Handle("/", http.FileServer(http.Dir("./milestone1")))
	// http.HandleFunc("/", handler) // todo: learn about passing functions into functions. Are they copied or passed by reference?
	http.ListenAndServe(":8080", nil) // todo: learn about second arg
}
