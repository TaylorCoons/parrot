package main

import (
	"fmt"
	"net/http"
	"net/url"
	"parrot/api/src/server"
)

func handler(w http.ResponseWriter, r *http.Request, u url.URL, p server.PathParams) {
	fmt.Println("I'm Called")
}
func main() {
	var routes = []server.Route{
		{Method: "GET", Path: "path", Handler: handler},
	}
	server := server.Server{Routes: routes}
	http.ListenAndServe(":8080", server)
}
