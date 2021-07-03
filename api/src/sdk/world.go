package sdk

import (
	"github.com/jameycribbs/hare"
)

type World struct {
	Name string `json:"name"`
}

const (
	table = "World"
)

func CreateWorld(c *hare.Database, name string) error {
	r := WorldRecord{Name: name}
	_, err := c.Insert(table, &r)
	if err != nil {
		return err
	}
	return nil
}
