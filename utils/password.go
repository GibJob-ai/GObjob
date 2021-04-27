package utils

import (
	"golang.org/x/crypto/bcrypt"
	"log"
)

// HashPass: hashes a password with a salt
// params: password string
// returns: string, the hashed password
func HashPass(pass string) string {
	bytePass := []byte(pass)
	hash, err := bcrypt.GenerateFromPassword(bytePass, bcrypt.DefaultCost)
	if err != nil {
		log.Printf("Hash error: %s", err)
		return ""
	}

	return string(hash)
}

// CompareHashPass: compares a hash and a password
// params: password string, hash string
// returns: bool, true if they match, false otherwise
func CompareHashPass(pass string, hash string) bool {
	byteHash := []byte(hash)
	bytePass := []byte(pass)

	// returns an err if false
	err := bcrypt.CompareHashAndPassword(byteHash, bytePass)

	if err != nil {
		return false
	}

	return true
}
