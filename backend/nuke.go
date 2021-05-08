package migrations

import (
	"github.com/GibJob-ai/GObjob/db"
	"github.com/GibJob-ai/GObjob/model"
)

// delete the tables
func Nuke(db *db.DB) {
	db.DropTableIfExists(&model.User{})
	db.DropTableIfExists(&model.File{})
}
