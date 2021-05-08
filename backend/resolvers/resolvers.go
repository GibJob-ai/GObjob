package resolvers

import (
	"github.com/GibJob-ai/GObjob/db"
)

// Resolvers including query and mutation
type Resolvers struct {
	*db.DB
}
