package milestone1

// todo: learn what are packages
// todo: learn named imports
// todo: learn about passing functions into functions. Are they copied or passed by reference?
// todo: learn http.Error status int
// todo: learn about second arg for ListenAndServe

import "net/http"


func Serve() {
  http.HandleFunc("/", RootHandler)
  http.Handle("/static", staticHandler)
	http.HandleFunc("/chat", ChatHandler)

	http.ListenAndServe(":8080", nil)
}
