package db

import (
	// "os"
	"fmt"
	"log"

	"github.com/jinzhu/gorm"

	"github.com/GibJob-ai/GObjob/config"
	// gorm postgres dialect
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// DB *grom.DB
type DB struct {
	*gorm.DB
}

// ConnectDB : connecting DB
func ConnectDB() (*DB, error) {

	// unix socket connection to cloud sql proxy
	// if config.CONFIG.Mode == "prod"{
		// dboptions := fmt.Sprintf(
			// "host=%s user=%s dbname=%s password=%s sslmode=%s",
			// config.CONFIG.DB.Host,
			// config.CONFIG.DB.User,
			// config.CONFIG.DB.DBName,
			// config.CONFIG.DB.Password,
			// config.CONFIG.DB.SSLMode,
		// )
	// }
	// local database for dev and test modes
	dboptions := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		config.CONFIG.DB.Host,
		config.CONFIG.DB.Port,
		config.CONFIG.DB.User,
		config.CONFIG.DB.DBName,
		config.CONFIG.DB.Password,
		config.CONFIG.DB.SSLMode,
	)

	// connect to db
	db, err := gorm.Open("postgres", dboptions)

	if err != nil {
		log.Fatalf("Couldnt connect to db: %s", err)
	}

	return &DB{db}, nil
}
