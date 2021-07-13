package routes

import (
	"encoding/json"
	"net/http"
	"parrot/api/src/connector"
	"parrot/api/src/sdk/world"
	"parrot/api/src/server"
)

var Routes []server.Route = []server.Route{
	{Method: "POST", Path: "/world", Handler: CreateWorldHandler},
	{Method: "GET", Path: "/world", Handler: GetWorldsHandler},
}

func CreateWorldHandler(w http.ResponseWriter, r *http.Request, p server.PathParams) {
	c := connector.GetConnector()
	err := world.CreateWorld(c, "MyWorld")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func GetWorldsHandler(w http.ResponseWriter, r *http.Request, p server.PathParams) {
	c := connector.GetConnector()
	worlds, err := world.GetWorlds(c)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	json.NewEncoder(w).Encode(worlds)
}
