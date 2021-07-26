package coord

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

type BiomeType string

const (
	Plains         BiomeType = "plains"
	Forest         BiomeType = "forest"
	Jungle         BiomeType = "jungle"
	Mountains      BiomeType = "mountains"
	Desert         BiomeType = "desert"
	Taiga          BiomeType = "taiga"
	SnowyTunda     BiomeType = "snowy tundra"
	IceSpikes      BiomeType = "ice spikes"
	Swamp          BiomeType = "swamp"
	Savannah       BiomeType = "savannah"
	Badlands       BiomeType = "badlands"
	Beach          BiomeType = "beach"
	StoneShore     BiomeType = "stone shore"
	River          BiomeType = "river"
	Ocean          BiomeType = "ocean"
	MushroomIsland BiomeType = "mushroom island"
	BasaltDelta    BiomeType = "basalt delta"
	CrimpsonForest BiomeType = "crimpson forest"
	NetherWastes   BiomeType = "nether wastes"
	SoulSandValley BiomeType = "soul sand valley"
	WarpedForest   BiomeType = "warped forest"
)

type CoordRecord struct {
	ID          int            `json:"id"`
	World       string         `json:"world"`
	Created     int            `json:"created"`
	Updated     int            `json:"updated"`
	X           int            `json:"x"`
	Y           int            `json:"y"`
	Z           int            `json:"z"`
	Realm       *RealmType     `json:"realm"`
	Structure   *StructureType `json:"structure"`
	Biome       *BiomeType     `json:"biome"`
	Description *string        `json:"description"`
}

// GetID returns the record id.
// This method is used internally by Hare.
// You need to add this method to each one of
// your models.
func (c *CoordRecord) GetID() int {
	return c.ID
}

// SetID takes an id. This method is used
// internally by Hare.
// You need to add this method to each one of
// your models.
func (c *CoordRecord) SetID(id int) {
	c.ID = id
}

// AfterFind is a callback that is run by Hare after
// a record is found.
// You need to add this method to each one of your
// models.
func (c *CoordRecord) AfterFind(db *hare.Database) error {
	// IMPORTANT!!!  These two lines of code are necessary in your AfterFind
	//               in order for the Find method to work correctly!
	*c = CoordRecord(*c)

	return nil
}
