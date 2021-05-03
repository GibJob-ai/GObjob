package migrations

import (
	"github.com/GibJob-ai/GObjob/db"
	"github.com/GibJob-ai/GObjob/model"
)

func migrate_2(d *db.DB) {
	// load the config
	d.DropTableIfExists(&model.User{}) // temp, this will be replaced with an actual new model
	d.CreateTable(&model.User{}) // temp, this will be replaced with an actual new model
}
