package routes

import (
	"net/http"
	"parrot/api/src/connector"
	sdk "parrot/api/src/sdk/world"
	"parrot/api/src/server"
)

var Routes []server.Route = []server.Route{
	{Method: "POST", Path: "/world", Handler: CreateWorldHandler},
}

func CreateWorldHandler(w http.ResponseWriter, r *http.Request, p server.PathParams) {
	c := connector.GetConnector()
	sdk.CreateWorld(c, "MyWorld")
}
