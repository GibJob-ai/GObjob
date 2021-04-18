package db

import (
	"os"
	"github.com/jinzhu/gorm"
  "fmt"

	// gorm postgres dialect
  "github.com/jinzhu/gorm/dialects/postgres"
)

// DB *grom.DB
type DB struct {
	*gorm.DB
}

// ConnectDB : connecting DB
func ConnectDB() (*DB, error) {

	db, err := gorm.Open("postgres", "host=localhost port=5432 user=gibjob dbname=gibjob password=gibjob sslmode=disable")

	if err != nil {
		panic(err)
	}

	return &DB{db}, nil
}
