package util

import "golang.org/x/crypto/bcrypt"

// Hash generates a hashed version of the provided password using bcrypt.
// It employs a cost factor of 14 to ensure a strong hash.
// The function returns the hashed password as a string and any potential error encountered during the hashing process.
func Hash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
