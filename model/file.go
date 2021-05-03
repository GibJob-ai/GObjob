package model

import (
	"github.com/jinzhu/gorm"
)

var supportedFiletypes = []string {"pdf", 'png', 'jpeg'} // possibly svg later for custom icon upload
var documentTypes = []string {'resume_pdf', 'resume_thumb', 'profile_picture'} // possibly svg later for custom icon upload

// File : Model with injected fields `ID`, `CreatedAt`, `UpdatedAt`
type File struct {
	Model
	Url     string `gorm:"type:varchar(160);not null"`
	FileType  string `gorm:"not null;check:filetype in ('test')"` // eg. PDF, PNG, etc.
	// DocumentType string // eg. RESUME, 
}
