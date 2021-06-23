package main

import (
	"fmt"
	"net/http"
	"parrot/api/src/server"
)

func handler(w http.ResponseWriter, r *http.Request, p server.PathParams) {
	fmt.Println("I'm Called")
}
func main() {
	var routes = []server.Route{
		{Method: "GET", Path: "/path", Handler: handler},
	}
	compiledRoutes := server.CompileRoutes(routes)
	server := server.Server{CompiledRoutes: compiledRoutes}
	http.ListenAndServe(":8080", server)
}
