package resolvers

import (
	"fmt"
	"log"

	"github.com/GibJob-ai/GObjob/model"
	"github.com/GibJob-ai/GObjob/utils"
)

// SignUp mutation creates user
func (r *Resolvers) SignUp(args signUpMutationArgs) (*UserResponse, error) {

	// hash the password
	hashPassword, err := utils.HashPass(args.Password)
	if err != nil {
		log.Printf("ERROR, couldnt hash password: %#v", err)
		return nil, err
	}

	newUser := model.User{Email: args.Email, Password: hashPassword, Username: args.Username}

	// if user already exists return an error of UserExists
	if !r.DB.Where("email = ?", args.Email).First(&model.User{}).RecordNotFound() {
		return nil, &userExistsError{Code: "UserExists", Message: "User already signed up"}
	}

	r.DB.Create(&newUser)

	return &UserResponse{u: &newUser}, nil
}

// arguements
type signUpMutationArgs struct {
	Email    string
	Password string
	Username string
}

// user exists error
type userExistsError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func (e userExistsError) Error() string {
	return fmt.Sprintf("error [%s]: %s", e.Code, e.Message)
}

func (e userExistsError) Extensions() map[string]interface{} {
	return map[string]interface{}{
		"code":    e.Code,
		"message": e.Message,
	}
}
