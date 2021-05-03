package migrations

import (
	"github.com/GibJob-ai/GObjob/db"
	"github.com/GibJob-ai/GObjob/model"
)

func migrate_1(d *db.DB) {
	// load the config
	d.DropTableIfExists(&model.User{})
	d.CreateTable(&model.User{})
}
