import (
	"log"
	"github.com/GibJob-ai/GObjob/db"
	"github.com/GibJob-ai/GObjob/model"
	"github.com/GibJob-ai/GObjob/config"
)

func main() {
	// load the config
	config.Load()
	d, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}

	defer d.Close()

	migrate_1_add_user(d)
	migrate_2_add_files(d)
}
