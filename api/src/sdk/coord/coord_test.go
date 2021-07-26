package coord

import (
	"fmt"
	"os"
	"parrot/api/src/connector"
	"path/filepath"
	"testing"
)

var testRealm = Overworld
var testStructure = DesertTemple
var testBiome = Desert
var testDescription = "Test Coordinate"

var testCoord Coord = Coord{
	X:           -1,
	Y:           1,
	Z:           4234193,
	Realm:       &testRealm,
	Structure:   &testStructure,
	Biome:       &testBiome,
	Description: &testDescription,
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

func TestGetCoords(t *testing.T) {
	c := connector.New(getFsPath("Coord"))
	defer c.Close()
	defer os.RemoveAll(getFsPath("Coord"))
	err := CreateCoord(c, "TestWorld", testCoord)
	if err != nil {
		fmt.Println(err)
		t.Error("Failed to create coord")
	}
	coords, err := GetCoords(c, "TestWorld")
	if err != nil {
		fmt.Println(err)
		t.Error("Failed to get coords")
	}
	if len(coords) != 1 {
		t.Error("Length of coords does not match")
	}
}

func TestGetCoord(t *testing.T) {
	c := connector.New(getFsPath("Coord"))
	defer c.Close()
	defer os.RemoveAll(getFsPath("Coord"))
	err := CreateCoord(c, "TestWorld", testCoord)
	if err != nil {
		fmt.Println(err)
		t.Error("Failed to create coord")
	}
	coords, err := GetCoords(c, "TestWorld")
	if err != nil {
		fmt.Println(err)
		t.Error("Failed to get coords")
	}
	_, err = GetCoord(c, "TestWorld", *coords[0].ID)
	if err != nil {
		fmt.Println(err)
		t.Error("Failed to get coord")
	}
}
