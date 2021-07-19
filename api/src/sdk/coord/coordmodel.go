package world

import (
	"github.com/jameycribbs/hare"
)

type RealmType string

const (
	Overworld RealmType = "overworld"
	Nether    RealmType = "nether"
	End       RealmType = "end"
)

type StructureType string

const (
	Mineshaft        StructureType = "mineshaft"
	Stronghold       StructureType = "stronghold"
	Dungeon          StructureType = "dungeon"
	DesertWell       StructureType = "desert well"
	Fossil           StructureType = "fossil"
	Village          StructureType = "village"
	DesertTemple     StructureType = "desert temple"
	WitchHut         StructureType = "witch hut"
	OceanMonument    StructureType = "ocean monument"
	Igloo            StructureType = "igloo"
	WoodlandMansion  StructureType = "woodland mansion"
	NetherFortress   StructureType = "nether fortress"
	BastionRemnant   StructureType = "bastion remntant"
	RuinedPortal     StructureType = "ruined portal"
	ObsidianPillar   StructureType = "obsidian pillar"
	EndFountain      StructureType = "end fountain"
	EndGatewayPortal StructureType = "end gateway portal"
	EndCity          StructureType = "end city"
	EndShip          StructureType = "end ship"
)

type WorldRecord struct {
	// Required field!!!
	ID        int           `json:"id"`
	Created   int           `json:"created"`
	Updated   int           `json:"updated"`
	X         int           `json:"x"`
	Y         int           `json:"y"`
	Z         int           `json:"z"`
	Realm     RealmType     `json:"realm"`
	Structure StructureType `json:"StructureType"`
}

// GetID returns the record id.
// This method is used internally by Hare.
// You need to add this method to each one of
// your models.
func (w *WorldRecord) GetID() int {
	return w.ID
}

// SetID takes an id. This method is used
// internally by Hare.
// You need to add this method to each one of
// your models.
func (w *WorldRecord) SetID(id int) {
	w.ID = id
}

// AfterFind is a callback that is run by Hare after
// a record is found.
// You need to add this method to each one of your
// models.
func (w *WorldRecord) AfterFind(db *hare.Database) error {
	// IMPORTANT!!!  These two lines of code are necessary in your AfterFind
	//               in order for the Find method to work correctly!
	*w = WorldRecord(*w)

	return nil
}
