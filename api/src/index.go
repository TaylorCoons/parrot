package main

import (
	"net/http"
	"parrot/api/src/connector"
	"parrot/api/src/routes"
	"parrot/api/src/server"
)

func main() {
	compiledRoutes := server.CompileRoutes(routes.Routes)
	server := server.Server{CompiledRoutes: compiledRoutes}
	c := connector.New()
	connector.SetConnector(c)
	err := http.ListenAndServe(":8080", server)
	if err != nil {
		panic(err)
	}
}
