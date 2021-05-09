package model

import "github.com/google/uuid"

var supportedFiletypes = []string{"pdf", "png", "jpeg"}                       // possibly svg later for custom icon upload
var documentTypes = []string{"resume_pdf", "resume_thumb", "profile_picture"} // possibly svg later for custom icon upload

// File : Model with injected fields `ID`, `CreatedAt`, `UpdatedAt`
type File struct {
	Model
	Url string `gorm:"type:varchar(160);not null"`

	// TODO someone figure out how to make this an enum?
	FileType string `gorm:"not null;` // eg. pdf, png, etc.

	// TODO someone figure out how to make this an enum?
	DocumentType string `gorm:"not null;` // eg. resume_pdf, profile_picture,

	Name string `gorm:"not null;` // eg. 'my resume 123'

	UserID uuid.UUID // foreign key
}

// FileType TODO add enums to File
type FileType int
const (
	pdf FileType = iota
	png
	jpeg
)

type DocumentType int
// This can probably be broken up even more/should probably have resume uploads/pfp uploads separated
const(
	resume_pdf DocumentType = iota
	profile_picture
)
