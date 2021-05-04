package main

import (
	"flag"
	"log"

	"github.com/gin-gonic/gin"

	"github.com/GibJob-ai/GObjob/config"
	"github.com/GibJob-ai/GObjob/db"
	"github.com/GibJob-ai/GObjob/migrations"
	"github.com/GibJob-ai/GObjob/server"
)

func main() {
	migrate_only := flag.Bool("migrate_only-only", false, "migrate_only db but dont serve")
	flag.Parse()

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

	db, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}

	defer db.Close()

	migrations.Migrate(db)
	if !*migrate_only {
		server.Serve(db)
	}
}
