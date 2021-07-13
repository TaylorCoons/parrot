package main

import (
	"net/http"
	"os"
	"parrot/api/src/connector"
	"parrot/api/src/routes"
	"parrot/api/src/server"
	"path/filepath"
)

func main() {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		panic("Failed to stat root path")
	}
	fsPath := filepath.Join(dir, "./filestore")
	compiledRoutes := server.CompileRoutes(routes.Routes)
	server := server.Server{CompiledRoutes: compiledRoutes}
	c := connector.New(fsPath)
	defer c.Close()
	connector.SetConnector(c)
	err = http.ListenAndServe(":8080", server)
	if err != nil {
		panic(err)
	}
}
