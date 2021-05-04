package resolvers

import (
	"context"

	"github.com/GibJob-ai/GObjob/model"
	"github.com/GibJob-ai/GObjob/handler"
)

// GetUser resolver
func (r *Resolvers) GetUser(ctx context.Context) (*GetUserResponse, error) {
	userID := ctx.Value(handler.ContextKey("userID"))

	if userID == nil {
		msg := "Not Authorized"
		return &GetUserResponse{Status: false, Msg: &msg, User: nil}, nil
	}

	user := model.User{}
	if err := r.DB.First(&user, userID).Error; err != nil {
		msg := "Not found"
		return &GetUserResponse{Status: false, Msg: &msg, User: nil}, nil
	}
	return &GetUserResponse{Status: true, Msg: nil, User: &UserResponse{u: &user}}, nil
}

// GetUserResponse is the response type
type GetUserResponse struct {
	Status bool
	Msg    *string
	User   *UserResponse
}

// Ok for GetUserResponse
func (r *GetUserResponse) Ok() bool {
	return r.Status
}

// Error for GetUserResponse
func (r *GetUserResponse) Error() *string {
	return r.Msg
}
