package connector

import (
	"os"
	"path/filepath"
	"testing"
)

func getFsPath() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		panic("Failed to stat root path")
	}
	return filepath.Join(dir, "./connector_test_filestore")
}

func TestNew(t *testing.T) {
	fsPath := getFsPath()
	c := New(getFsPath())
	if c == nil {
		t.Error("Nil connector returned")
	}
	defer os.Remove(fsPath)
	_, err := os.Stat(fsPath)
	if err != nil {
		t.Error("New did not create the filestore directory")
	}
}

func TestSetGetConnector(t *testing.T) {
	fsPath := getFsPath()
	c := New(getFsPath())
	defer os.Remove(fsPath)
	SetConnector(c)
	if c != GetConnector() {
		t.Error("Set and get connector failed")
	}
}

func TestSetGetConnector_Panics(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("Get did not panic on nil connector")
		}
	}()
	SetConnector(nil)
	GetConnector()
}
