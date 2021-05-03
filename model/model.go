
package model

import (
	"time"
	"github.com/jinzhu/gorm"
)

var _ := gorm; // make linter happy by using gorm TODO needed??

// Model : gorm.Model definition
type Model struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

