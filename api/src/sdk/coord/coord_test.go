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
	err := CreateCoord(c, "TestWorld", testCoord)
	if err != nil {
		fmt.Println(err)
		t.Error("Failed to create coord")
	}
}

func TestCreateCoord(t *testing.T) {
	c := connector.New(getFsPath("Coord"))
	defer c.Close()
	defer os.RemoveAll(getFsPath("Coord"))
	createCoordHelper(t, c, "TestWorld", testCoord)
}

func TestGetCoords(t *testing.T) {
	c := connector.New(getFsPath("Coord"))
	defer c.Close()
	defer os.RemoveAll(getFsPath("Coord"))
	createCoordHelper(t, c, "TestWorld", testCoord)
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
	createCoordHelper(t, c, "TestWorld", testCoord)
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

func TestUpdateCoord(t *testing.T) {
	c := connector.New(getFsPath("Coord"))
	defer c.Close()
	defer os.RemoveAll(getFsPath("Coord"))
	createCoordHelper(t, c, "TestWorld", testCoord)
	coords, err := GetCoords(c, "TestWorld")
	if err != nil {
		fmt.Println(err)
		t.Error("Failed to get coords")
	}
	updateDescription := "I have been updated"
	coords[0].Description = &updateDescription
	err = UpdateCoord(c, "TestWorld", *coords[0].ID, coords[0])
	if err != nil {
		fmt.Println(err)
		t.Error("Failed to update coord")
	}
	coord, err := GetCoord(c, "TestWorld", *coords[0].ID)
	if err != nil {
		fmt.Println(err)
		t.Error("Failed to get coord after update")
	}
	if *coord.Description != updateDescription {
		t.Error("Updated description does not match")
	}
}

func TestDeleteCoord(t *testing.T) {
	c := connector.New(getFsPath("Coord"))
	defer c.Close()
	defer os.RemoveAll(getFsPath("Coord"))
	createCoordHelper(t, c, "TestWorld", testCoord)
	coords, err := GetCoords(c, "TestWorld")
	if err != nil {
		fmt.Println(err)
		t.Error("Failed to get coords")
	}
	err = DeleteCoord(c, "TestWorld", *coords[0].ID)
	if err != nil {
		fmt.Println(err)
		t.Error("Failed to delete coord")
	}
	coords, err = GetCoords(c, "TestWorld")
	if err != nil {
		fmt.Println(err)
		t.Error("Failed to get coords")
	}
	if len(coords) != 0 {
		t.Error("Coord failed to delete")
	}

}
