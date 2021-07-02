package sdk

import (
	"fmt"

	"github.com/jameycribbs/hare"
)

type World struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func CreateWorld(c *hare.Database, name string) {
	fmt.Println("Creating wrold")
}
