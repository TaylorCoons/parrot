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

func GetWorlds(c *hare.Database) ([]string, error) {
	ids, err := c.IDs(table)
	if err != nil {
		return nil, err
	}
	worlds := make([]string, len(ids))
	for i, id := range ids {
		r := WorldRecord{}
		err = c.Find(table, id, &r)
		if err != nil {
			return nil, err
		}
		worlds[i] = r.Name
	}
	return worlds, nil
}

func DeleteWorld(c *hare.Database, name string) error {
	ids, err := c.IDs(table)
	if err != nil {
		return err
	}
	for _, id := range ids {
		r := WorldRecord{}
		err = c.Find(table, id, &r)
		if err != nil {
			return err
		}
		if r.Name == name {
			err := c.Delete(table, id)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
