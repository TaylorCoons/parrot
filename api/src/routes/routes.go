package routes

import (
	"encoding/json"
	"net/http"
	"parrot/api/src/connector"
	"parrot/api/src/sdk/coord"
	"parrot/api/src/sdk/world"
	"strconv"

	server "github.com/TaylorCoons/gorouter"
)

var Routes []server.Route = []server.Route{
	{Method: "POST", Path: "/world", Handler: CreateWorldHandler},
	{Method: "GET", Path: "/world", Handler: GetWorldsHandler},
	{Method: "DELETE", Path: "/world/:world", Handler: DeleteWorldHandler},
	{Method: "DELETE", Path: "/world", Handler: DeleteWorldsHandler},
	{Method: "GET", Path: "/world/:world/coord", Handler: GetCoordsHandler},
	{Method: "POST", Path: "/world/:world/coord", Handler: CreateCoordHandler},
	{Method: "GET", Path: "/world/:world/coord/:coordId", Handler: GetCoordHandler},
	{Method: "PUT", Path: "/world/:world/coord/:coordId", Handler: UpdateCoordHandler},
	{Method: "DELETE", Path: "/world/:world/coord/:coordId", Handler: DeleteCoordHandler},
}

func CreateWorldHandler(w http.ResponseWriter, r *http.Request, p server.PathParams) {
	c := connector.GetConnector()
	// TODO: Don't hardcode worldname
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
	// TODO: catch error for when encoding fails
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

func GetCoordsHandler(w http.ResponseWriter, r *http.Request, p server.PathParams) {
	c := connector.GetConnector()
	coords, err := coord.GetCoords(c, p["world"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// TODO: catch error for when encoding fails
	json.NewEncoder(w).Encode(coords)
}

func CreateCoordHandler(w http.ResponseWriter, r *http.Request, p server.PathParams) {
	c := connector.GetConnector()
	coordData := coord.Coord{}
	err := json.NewDecoder(r.Body).Decode(&coordData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = coord.CreateCoord(c, p["world"], coordData)
	if err != nil {
		if _, ok := err.(*coord.InvalidCoordError); ok {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func GetCoordHandler(w http.ResponseWriter, r *http.Request, p server.PathParams) {
	c := connector.GetConnector()
	coordId, err := strconv.Atoi(p["coordId"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	coordData, err := coord.GetCoord(c, p["world"], coordId)
	if err != nil {
		if _, ok := err.(*coord.CoordNotExistError); ok {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// TODO: catch error for when encoding fails
	json.NewEncoder(w).Encode(coordData)
}

func UpdateCoordHandler(w http.ResponseWriter, r *http.Request, p server.PathParams) {
	c := connector.GetConnector()
	coordId, err := strconv.Atoi(p["coordId"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	coordData := coord.Coord{}
	err = json.NewDecoder(r.Body).Decode(&coordData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = coord.UpdateCoord(c, p["world"], coordId, coordData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func DeleteCoordHandler(w http.ResponseWriter, r *http.Request, p server.PathParams) {
	c := connector.GetConnector()
	coordId, err := strconv.Atoi(p["coordId"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = coord.DeleteCoord(c, p["world"], coordId)
	if err != nil {
		if _, ok := err.(*coord.CoordNotExistError); ok {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
