package model

import (
	"github.com/jinzhu/gorm"
)

// User : Model with injected fields `ID`, `CreatedAt`, `UpdatedAt`
type User struct {
	Model
	Email     string `gorm:"type:varchar(100);not null"`
	Password  string `gorm:"not null"`
	FirstName string `gorm:"type:varchar(50);not null"`
	LastName  string `gorm:"type:varchar(50);not null"`
	Bio       string
	Avatar    string
}
