package routes

import (
	"fmt"
	"net/http"
	"parrot/api/src/server"
)

func pongHandler(w http.ResponseWriter, r *http.Request, p server.PathParams) {
	fmt.Fprint(w, "pong\n")
}

var Routes []server.Route = []server.Route{
	{Method: "GET", Path: "/ping", Handler: pongHandler},
}
