package sdk

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
