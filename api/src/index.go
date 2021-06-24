package main

import (
	"net/http"
	"parrot/api/src/routes"
	"parrot/api/src/server"
)

func main() {
	compiledRoutes := server.CompileRoutes(routes.Routes)
	server := server.Server{CompiledRoutes: compiledRoutes}
	http.ListenAndServe(":8080", server)
}
