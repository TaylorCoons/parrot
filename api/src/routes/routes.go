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
	{Method: "DELETE", Path: "/world/:world", Handler: DeleteWorldHandler},
	{Method: "DELETE", Path: "/world", Handler: DeleteWorldsHandler},
}

func CreateWorldHandler(w http.ResponseWriter, r *http.Request, p server.PathParams) {
	c := connector.GetConnector()
	err := world.CreateWorld(c, "MyWorld")
	if err != nil {
		if _, ok := err.(*world.DuplicateWorldError); ok {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func GetWorldsHandler(w http.ResponseWriter, r *http.Request, p server.PathParams) {
	c := connector.GetConnector()
	worlds, err := world.GetWorlds(c)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(worlds)
}

func DeleteWorldHandler(w http.ResponseWriter, r *http.Request, p server.PathParams) {
	c := connector.GetConnector()
	err := world.DeleteWorld(c, p["world"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func DeleteWorldsHandler(w http.ResponseWriter, r *http.Request, p server.PathParams) {
	c := connector.GetConnector()
	err := world.DeleteWorlds(c)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
