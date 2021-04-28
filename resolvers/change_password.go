package resolvers

import (
	"context"
	"log"

	"github.com/GibJob-ai/GObjob/handler"
	"github.com/GibJob-ai/GObjob/model"
	"github.com/GibJob-ai/GObjob/utils"
)

// ChangePassword mutation change password
func (r *Resolvers) ChangePassword(ctx context.Context, args changePasswordMutationArgs) (*ChangePasswordResponse, error) {
	userID := ctx.Value(handler.ContextKey("userID"))

	if userID == nil {
		msg := "Not Authorized"
		return &ChangePasswordResponse{Status: false, Msg: &msg, User: nil}, nil
	}
	user := model.User{}

	if err := r.DB.First(&user, userID).Error; err != nil {
		msg := "Not existing user"
		return &ChangePasswordResponse{Status: false, Msg: &msg, User: nil}, nil
	}

	hashPassword, err := utils.HashPass(args.Password)
	if err != nil {
		log.Printf("ERROR, couldnt hash password: %#v", err)
		return nil, err
	}

	user.Password = hashPassword

	r.DB.Save(&user)
	return &ChangePasswordResponse{Status: true, Msg: nil, User: &UserResponse{u: &user}}, nil
}

type changePasswordMutationArgs struct {
	Password string
}

// ChangePasswordResponse is the response type
type ChangePasswordResponse struct {
	Status bool
	Msg    *string
	User   *UserResponse
}

// Ok for ChangePasswordResponse
func (r *ChangePasswordResponse) Ok() bool {
	return r.Status
}

// Error for ChangePasswordResponse
func (r *ChangePasswordResponse) Error() *string {
	return r.Msg
}
