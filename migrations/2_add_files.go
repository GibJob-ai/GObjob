package migrations

import (
	"github.com/GibJob-ai/GObjob/db"
	"github.com/GibJob-ai/GObjob/model"
)

func migrate_2_add_files(db *db.DB) {
	db.AutoMigrate(&model.File{})
}
