package migrations

import (
	"github.com/GibJob-ai/GObjob/db"
	"github.com/GibJob-ai/GObjob/model"
)

func migrate_1_add_user(d *db.DB) {
	// load the config
	d.DropTableIfExists(&model.User{})
	d.CreateTable(&model.User{})
}
