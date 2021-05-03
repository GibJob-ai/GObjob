package migrations

import (
	"github.com/GibJob-ai/GObjob/db"
	"github.com/GibJob-ai/GObjob/model"
)

func migrate_2_add_files(d *db.DB) {
	d.DropTableIfExists(&model.File{})
	d.CreateTable(&model.File{})
}
