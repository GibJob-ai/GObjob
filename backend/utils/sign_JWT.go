package utils

import (
	jwt "github.com/dgrijalva/jwt-go"
	"time"
)

// SignJWT : func to generate JWT
func SignJWT(userID *string) (*string, error) {
	claims := jwt.MapClaims{}
	claims["userID"] = *userID
	claims["exp"] = time.Now().Add(time.Second * 30 * 24 * 60 * 60)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte("my_secret"))

	return &tokenString, err
}
