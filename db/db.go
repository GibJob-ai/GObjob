package db

import (
	// "os"
	"fmt"

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

	// format them all to a string
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
	//db, err := gorm.Open("postgres", "host=localhost port=5432 user=gibjob dbname=gibjob password=gibjob sslmode=disable")

	if err != nil {
		panic(err)
	}

	return &DB{db}, nil
}
