package resolvers

import (
	"fmt"
	"github.com/GibJob-ai/GObjob/model"
	"github.com/google/uuid"
)

// SaveFileToDB mutation to create file
func (r *Resolvers) SaveFileToDB(args saveFileArgs) (*FileResponse, error){
	if !r.DB.Where("UUID = ?", args.UUID).First(&model.File{}).RecordNotFound() {
		return nil, &fileCreationError{Code:"FileExists", Message: "File already exists within DB"}
	}

	newFile := model.File{FileType: "pdf", DocumentType: "resume_pdf", Name: args.Filename, UserID: args.UUID}
	result := r.DB.Create(&newFile)
	if result.Error != nil {
		return nil, &fileCreationError{Code: "DBErr", Message: "Database had an error"}
	}
	// TODO basically we want to be able to access this function from save file
	return &FileResponse{f: &newFile}, nil

}

type saveFileArgs struct {
	UUID uuid.UUID
	Filename string
}

// file exists error
type fileCreationError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func (e fileCreationError) Error() string {
	return fmt.Sprintf("error [%s]: %s", e.Code, e.Message)
}

func (e fileCreationError) Extensions() map[string]interface{} {
	return map[string]interface{}{
		"code":    e.Code,
		"message": e.Message,
	}
}