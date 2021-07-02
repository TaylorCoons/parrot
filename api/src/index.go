package main

import (
	"net/http"
	"parrot/api/src/routes"
	"parrot/api/src/server"
)

func main() {
	compiledRoutes := server.CompileRoutes(routes.Routes)
	server := server.Server{CompiledRoutes: compiledRoutes}
	err := http.ListenAndServe(":8080", server)
	if err != nil {
		panic(err)
	}

}
