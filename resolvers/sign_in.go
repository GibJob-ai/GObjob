package resolvers

import (
	"strconv"

	"github.com/GibJob-ai/GObjob/model"
	"github.com/GibJob-ai/GObjob/utils"
)

// SignIn mutation creates user
func (r *Resolvers) SignIn(args signInMutationArgs) (*SignInResponse, error) {
	user := model.User{}

	r.DB.Where("email = ?", args.Email).First(&user)

	if user.ID == 0 {
		msg := "Not Sign up yet"
		return &SignInResponse{Status: false, Msg: &msg, Token: nil}, nil
	}

	if !utils.CompareHashPass(args.Password, user.Password) {
		msg := "Password is not correct"
		return &SignInResponse{Status: false, Msg: &msg, Token: nil}, nil
	}

	userIDString := strconv.Itoa(int(user.ID))
	tokenString, err := utils.SignJWT(&userIDString)
	if err != nil {
		msg := "Error in generating JWT"
		return &SignInResponse{Status: false, Msg: &msg, Token: nil}, nil
	}

	return &SignInResponse{Status: true, Msg: nil, Token: tokenString}, nil
}

type signInMutationArgs struct {
	Email    string
	Password string
}

// SignInResponse is the response type
type SignInResponse struct {
	Status bool
	Msg    *string
	Token  *string
}

// Ok for SignUpResponse
func (r *SignInResponse) Ok() bool {
	return r.Status
}

// Error for SignUpResponse
func (r *SignInResponse) Error() *string {
	return r.Msg
}
