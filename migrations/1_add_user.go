package migrations

import (
	"github.com/GibJob-ai/GObjob/db"
	"github.com/GibJob-ai/GObjob/model"
)

func migrate_1_add_user(db *db.DB) {
	db.AutoMigrate(&model.User{})
}
