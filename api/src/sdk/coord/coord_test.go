
package coord

import (
	"fmt"
	"os"
	"parrot/api/src/connector"
	"path/filepath"
	"testing"
)

testCoord := Coord{
	X:           -1,
	Y:           1,
	Z:           4234193,
	Realm:       Overworld,
	Structure:   DesertTemple,
	Biome:       Desert,
	Description: "Test Coordinate",
}

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

func TestCreateTableIfNotExists(t *testing.T) {
	c := connector.New("./test_filestore")
	defer c.Close()
	if c.TableExists(table) {
		t.Error("Table exists before test started")
	}
	createTableIfNotExists(c)
	if !c.TableExists(table) {
		t.Error("Table was not created")
	}
}

func TestCreateCoord(t *testing.T) {
	c := connector.New(getFsPath("Coord"))
	defer c.Close()
	defer os.RemoveAll(getFsPath("Coord"))
	err := CreateCoord(c, "TestWorld", testCoord)
	if err != nil {
		fmt.Println(err)
		t.Error("Failed to create coord")
	}
}

func TestDeleteWorlds(t *testing.T) {
	c := connector.New(getFsPath("Coord"))
	defer c.Close()
	defer os.RemoveAll(getFsPath("Coord"))
	err := CreateCoord(c, "TestWorld", testCoord)
	if err != nil {
		fmt.Println(err)
		t.Error("Failed to create world")
	}
	err = CreateWorld(c, "TestWorld2")
	if err != nil {
		fmt.Println(err)
		t.Error("Failed to create second world")
	}
	err = DeleteWorlds(c)
	if err != nil {
		fmt.Println(err)
		t.Error("Failed to delete worlds")
	}
}

func sContains(s []string, m string) bool {
	for _, v := range s {
		if v == m {
			return true
		}
	}
	return false
}

func TestGetWorlds(t *testing.T) {
	c := connector.New(getFsPath("World"))
	defer c.Close()
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
	if !sContains(worlds, "World_1") {
		t.Error("First returned world does not match")
	}
	if !sContains(worlds, "World_2") {
		t.Error("Second returned world does not match")
	}
}

func TestDeleteWorld(t *testing.T) {
	c := connector.New(getFsPath("World"))
	defer c.Close()
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
