import (
	"log"
	"github.com/GibJob-ai/GObjob/db"
	"github.com/GibJob-ai/GObjob/model"
	"github.com/GibJob-ai/GObjob/config"
	"github.com/gin-gonic/gin"
)

func main() {
	// load the config
	config.Load()

	// if the config is dev mode set server to dev mode
	// same with other modes
	if config.CONFIG.Mode == "dev" {
		gin.SetMode(gin.DebugMode)
	} else if config.CONFIG.Mode == "prod" {
		gin.SetMode(gin.ReleaseMode)
	} else if config.CONFIG.Mode == "test" {
		gin.SetMode(gin.TestMode)
	} else {
		log.Fatalf("Unkown server mode given in config: %s", config.CONFIG.Mode)
	}
	d, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}

	defer d.Close()

	migrate_1(d)
	migrate_2(d)
}
