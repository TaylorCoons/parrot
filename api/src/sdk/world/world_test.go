package world

import (
	"fmt"
	"os"
	"parrot/api/src/connector"
	"path/filepath"
	"testing"
)

func getFsPath(table string) string {
	basePath := "./test_filestore"
	err := os.MkdirAll(basePath, 0777)
	if err != nil {
		panic(err)
	}
	fd, err := os.Create(filepath.Join(basePath, table+".json"))
	if err != nil {
		panic(err)
	}
	fd.Close()
	return basePath
}

func TestCreateWorld(t *testing.T) {
	c := connector.New(getFsPath("World"))
	defer os.RemoveAll(getFsPath("World"))
	err := CreateWorld(c, "TestWorld")
	if err != nil {
		fmt.Println(err)
		t.Error("Failed to create world")
	}
}

func TestCreateWorld_BadTable(t *testing.T) {
	c := connector.New(getFsPath("BadTable"))
	defer os.RemoveAll(getFsPath("BadTable"))
	err := CreateWorld(c, "TestWorld")
	if err == nil {
		t.Error("CreateWorld did not return error")
	}
}

func TestGetWorlds(t *testing.T) {
	c := connector.New(getFsPath("World"))
	defer os.RemoveAll(getFsPath("World"))
	err := CreateWorld(c, "World_1")
	if err != nil {
		t.Error("Failed to create world")
	}
	err = CreateWorld(c, "World_2")
	if err != nil {
		t.Error("Failed to create world")
	}
	worlds, err := GetWorlds(c)
	if err != nil {
		t.Error("Failed to create worlds")
	}
	if len(worlds) != 2 {
		t.Error("Returned worlds does not match")
	}
	if worlds[0] != "World_1" {
		t.Error("First returned world does not match")
	}
	if worlds[1] != "World_2" {
		t.Error("Second returned world does not match")
	}
}

func TestDeleteWorld(t *testing.T) {
	c := connector.New(getFsPath("World"))
	defer os.RemoveAll(getFsPath("World"))
	err := CreateWorld(c, "World_1")
	if err != nil {
		t.Error("Create world failed")
	}
	worlds, err := GetWorlds(c)
	if err != nil || len(worlds) != 1 {
		t.Error("Get worlds failed")
	}
	err = DeleteWorld(c, "World_1")
	if err != nil {
		t.Error("Delete world failed")
	}
	worlds, err = GetWorlds(c)
	if err != nil || len(worlds) != 0 {
		t.Error("Delete world returned success but did not remove record")
	}
}
