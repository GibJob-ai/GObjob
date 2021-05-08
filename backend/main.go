package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"github.com/GibJob-ai/GObjob/cmd"
	"github.com/GibJob-ai/GObjob/config"
	"github.com/GibJob-ai/GObjob/db"
	"github.com/GibJob-ai/GObjob/migrations"
	"github.com/GibJob-ai/GObjob/server"
)

func main() {
	cmdResults := cmd.Execute()

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

	if *cmdResults.Nuke {
		migrations.Nuke(db)
		migrations.Migrate(db)
		log.Print("Succesfully Recreated db")
	} else if *cmdResults.Migrate {
		migrations.Migrate(db)
		log.Print("Succesfully Migrated db")
	} else if *cmdResults.Migrate_serve {
		migrations.Migrate(db)
		log.Print("Succesfully Migrated db")
		server.Serve(db)
	} else {
		server.Serve(db)
	}

}
