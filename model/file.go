package model

import (
)

var supportedFiletypes = []string {"pdf", "png", "jpeg"} // possibly svg later for custom icon upload
var documentTypes = []string {"resume_pdf", "resume_thumb", "profile_picture"} // possibly svg later for custom icon upload

// File : Model with injected fields `ID`, `CreatedAt`, `UpdatedAt`
type File struct {
	Model
	Url     string `gorm:"type:varchar(160);not null"`
	FileType  string `gorm:"not null;` // eg. PDF, PNG, etc.
	//check:filetype in ("test")"
	// DocumentType string // eg. RESUME, 
}
