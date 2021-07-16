package connector

import (
	"os"

	"github.com/jameycribbs/hare"
	"github.com/jameycribbs/hare/datastores/disk"
)

var db *hare.Database

func New(fsPath string) *hare.Database {
	err := os.MkdirAll(fsPath, 0777)
	if err != nil {
		panic(err)
	}
	ds, err := disk.New(fsPath, ".json")
	if err != nil {
		panic(err)
	}
	db, err := hare.New(ds)
	if err != nil {
		panic(err)
	}
	return db
}

func SetConnector(d *hare.Database) {
	db = d
}

func GetConnector() *hare.Database {
	if db == nil {
		panic("connector is not initialized")
	}
	return db
}
