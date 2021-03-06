package coord

import (
	"fmt"
	"os"
	"parrot/api/src/connector"
	"path/filepath"
	"testing"

	"github.com/jameycribbs/hare"
)

var testX = -1
var testY = 1
var testZ = 4234193
var testRealm = Overworld
var testStructure = DesertTemple
var testBiome = Desert
var testDescription = "Test Coordinate"
var testWorld = "TestWorld"
var testDBPath = "Coord"

var testCoord Coord = Coord{
	X:           &testX,
	Y:           &testY,
	Z:           &testZ,
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

func createCoordHelper(t *testing.T, c *hare.Database, worldName string, coord Coord) {
	err := CreateCoord(c, testWorld, testCoord)
	if err != nil {
		fmt.Println(err)
		t.Error("Failed to create coord")
	}
}

func getCoordsHelper(t *testing.T, c *hare.Database, worldName string) []Coord {
	coords, err := GetCoords(c, worldName)
	if err != nil {
		fmt.Println(err)
		t.Error("Failed to get coords")
	}
	return coords
}

func getCoordHelper(t *testing.T, c *hare.Database, worldName string, coordId int) Coord {
	coord, err := GetCoord(c, worldName, coordId)
	if err != nil {
		fmt.Println(err)
		t.Error("Failed to get coord")
	}
	return coord
}

func updateCoordHelper(t *testing.T, c *hare.Database, worldName string, coordId int, coord Coord) {
	err := UpdateCoord(c, worldName, coordId, coord)
	if err != nil {
		fmt.Println(err)
		t.Error("Failed to update coord")
	}
}

func deleteCoordHelper(t *testing.T, c *hare.Database, worldName string, coordId int) {
	err := DeleteCoord(c, worldName, coordId)
	if err != nil {
		fmt.Println(err)
		t.Error("Failed to delete coord")
	}
}

func TestCreateCoord(t *testing.T) {
	c := connector.New(getFsPath(testDBPath))
	defer c.Close()
	defer os.RemoveAll(getFsPath(testDBPath))
	createCoordHelper(t, c, testWorld, testCoord)
}

func TestGetCoords(t *testing.T) {
	c := connector.New(getFsPath(testDBPath))
	defer c.Close()
	defer os.RemoveAll(getFsPath(testDBPath))
	createCoordHelper(t, c, testWorld, testCoord)
	coords := getCoordsHelper(t, c, testWorld)
	if len(coords) != 1 {
		t.Error("Length of coords does not match")
	}
}

func TestGetCoord(t *testing.T) {
	c := connector.New(getFsPath(testDBPath))
	defer c.Close()
	defer os.RemoveAll(getFsPath(testDBPath))
	createCoordHelper(t, c, testWorld, testCoord)
	coords := getCoordsHelper(t, c, testWorld)
	getCoordHelper(t, c, testWorld, *coords[0].ID)
}

func TestUpdateCoord(t *testing.T) {
	c := connector.New(getFsPath(testDBPath))
	defer c.Close()
	defer os.RemoveAll(getFsPath(testDBPath))
	createCoordHelper(t, c, testWorld, testCoord)
	coords := getCoordsHelper(t, c, testWorld)
	updateDescription := "I have been updated"
	coords[0].Description = &updateDescription
	updateCoordHelper(t, c, testWorld, *coords[0].ID, coords[0])
	coord := getCoordHelper(t, c, testWorld, *coords[0].ID)
	if *coord.Description != updateDescription {
		t.Error("Updated description does not match")
	}
}

func TestDeleteCoord(t *testing.T) {
	c := connector.New(getFsPath(testDBPath))
	defer c.Close()
	defer os.RemoveAll(getFsPath(testDBPath))
	createCoordHelper(t, c, testWorld, testCoord)
	coords := getCoordsHelper(t, c, testWorld)
	deleteCoordHelper(t, c, testWorld, *coords[0].ID)
	coords = getCoordsHelper(t, c, testWorld)
	if len(coords) != 0 {
		t.Error("Coord failed to delete")
	}
}
