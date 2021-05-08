package migrations

import (
	"github.com/GibJob-ai/GObjob/db"
)

// migrate the db
func Migrate(db *db.DB) {
	migrate_1_add_user(db)
	migrate_2_add_files(db)
}
