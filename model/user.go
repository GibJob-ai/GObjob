package model

import (
	"golang.org/x/crypto/bcrypt"
	"time"

	"github.com/GibJob-ai/GObjob/utils"
	"github.com/jinzhu/gorm"
	// gorm postgres dialect
	// "github.com/jinzhu/gorm/dialects/postgres"
)

// Model : gorm.Model definition
type Model struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

// User : Model with injected fields `ID`, `CreatedAt`, `UpdatedAt`
type User struct {
	gorm.Model
	Email     string `gorm:"type:varchar(100);not null"`
	Password  string `gorm:"not null"`
	FirstName string `gorm:"type:varchar(50);not null"`
	LastName  string `gorm:"type:varchar(50);not null"`
	Bio       string
	Avatar    string
}
