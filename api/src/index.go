package main

import (
	"fmt"
	"net/http"
	"os"
	"parrot/api/src/connector"
	"parrot/api/src/routes"
	"parrot/api/src/util"
	"path/filepath"

	server "github.com/TaylorCoons/gorouter"
)

func main() {
	startServer()
}

func startServer() {
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
	port := util.GetPort()
	bind := fmt.Sprintf(":%d", port)
	fmt.Println("Parrot is listening")
	err = http.ListenAndServe(bind, server)
	if err != nil {
		panic(err)
	}
}
