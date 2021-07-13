package world

import (
	"github.com/jameycribbs/hare"
)

type WorldRecord struct {
	// Required field!!!
	ID   int    `json:"id"`
	Name string `json:"name"`
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
