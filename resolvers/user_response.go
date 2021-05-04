package resolvers

import (
	"strconv"

	graphql "github.com/graph-gophers/graphql-go"

	"github.com/GibJob-ai/GObjob/model"
)

// UserResponse is the user response type
type UserResponse struct {
	u *model.User
}

// ID for UserResponse
func (r *UserResponse) ID() graphql.ID {
	id := strconv.Itoa(int(r.u.ID))
	return graphql.ID(id)
}

// Email for UserResponse
func (r *UserResponse) Email() string {
	return r.u.Email
}

// Password for UserResponse
func (r *UserResponse) Password() string {
	return r.u.Password
}

// Username for UserResponse
func (r *UserResponse) Username() string {
	return r.u.Username
}

// CreatedAt for UserResponse
func (r *UserResponse) CreatedAt() string {
	return r.u.CreatedAt.String()
}

// UpdatedAt for UserResponse
func (r *UserResponse) UpdatedAt() string {
	return r.u.UpdatedAt.String()
}

// Files for UserResponse
func (r *UserResponse) Files() []*fileResolver {
	var files []*fileResolver
	for _, file := range r.u.Files {
		files = append(files, &fileResolver{&file})
	}
	return files
}

type fileResolver struct {
	f *model.File
}

func (r *fileResolver) ID() graphql.ID {
	id := strconv.Itoa(int(r.f.ID))
	return graphql.ID(id)
}

func (r *fileResolver) Url() string {
	return r.f.Url
}

func (r *fileResolver) FileType() string {
	return r.f.FileType
}

func (r *fileResolver) DocumentType() string {
	return r.f.DocumentType
}

func (r *fileResolver) Name() string {
	return r.f.Name
}

func (r *fileResolver) CreatedAt() string {
	return r.f.CreatedAt.String()
}

func (r *fileResolver) UpdatedAt() string {
	return r.f.UpdatedAt.String()
}

func (r *fileResolver) User() int {
	return 1
}
