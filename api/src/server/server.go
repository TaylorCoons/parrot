package server

import (
	"fmt"
	"net/http"
)

type Server struct {
	CompiledRoutes []CompiledRoute
}

func (Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("I'm called")
}
