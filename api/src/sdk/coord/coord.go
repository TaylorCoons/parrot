package coord

import (
	"fmt"
	"time"

	"github.com/jameycribbs/hare"
)

type Coord struct {
	ID          *int           `json:"id"`
	Created     *int64         `json:"created"`
	Updated     *int64         `json:"updated"`
	X           int            `json:"x"`
	Y           int            `json:"y"`
	Z           int            `json:"z"`
	Realm       *RealmType     `json:"realm"`
	Structure   *StructureType `json:"structure"`
	Biome       *BiomeType     `json:"biome"`
	Description *string        `json:"description"`
}

const (
	table = "Coord"
)

type CoordNotExistError struct {
	CoordId int
}

func (cnee *CoordNotExistError) Error() string {
	return fmt.Sprintf("Coord: %d does not exist", cnee.CoordId)
}

func createTableIfNotExists(c *hare.Database) error {
	if !c.TableExists(table) {
		return c.CreateTable(table)
	}
	return nil
}

func CreateCoord(c *hare.Database, world string, coord Coord) error {
	err := createTableIfNotExists(c)
	if err != nil {
		return err
	}
	r := CoordRecord{
		World:       world,
		Created:     time.Now().Unix(),
		Updated:     time.Now().Unix(),
		X:           coord.X,
		Y:           coord.Y,
		Z:           coord.Z,
		Realm:       coord.Realm,
		Structure:   coord.Structure,
		Biome:       coord.Biome,
		Description: coord.Description,
	}
	_, err = c.Insert(table, &r)
	if err != nil {
		return err
	}
	return nil
}

func GetCoords(c *hare.Database, world string) ([]Coord, error) {
	err := createTableIfNotExists(c)
	if err != nil {
		return nil, err
	}
	ids, err := c.IDs(table)
	if err != nil {
		return nil, err
	}
	coords := make([]Coord, 0)
	for _, id := range ids {
		r := CoordRecord{}
		err = c.Find(table, id, &r)
		if err != nil {
			return nil, err
		}
		if r.World != world {
			continue
		}
		coord := coordRecordToCoord(r)
		coords = append(coords, coord)
	}
	return coords, nil
}

func GetCoord(c *hare.Database, world string, coordId int) (Coord, error) {
	err := createTableIfNotExists(c)
	if err != nil {
		return Coord{}, err
	}
	ids, err := c.IDs(table)
	if err != nil {
		return Coord{}, err
	}
	for _, id := range ids {
		if id == coordId {
			r := CoordRecord{}
			err = c.Find(table, id, &r)
			if err != nil {
				return Coord{}, err
			}
			return coordRecordToCoord(r), nil
		}
	}
	return Coord{}, &CoordNotExistError{coordId}
}

func coordRecordToCoord(r CoordRecord) Coord {
	return Coord{
		ID:          &r.ID,
		Created:     &r.Created,
		Updated:     &r.Updated,
		X:           r.X,
		Y:           r.Y,
		Z:           r.Z,
		Realm:       r.Realm,
		Structure:   r.Structure,
		Biome:       r.Biome,
		Description: r.Description,
	}
}
