package util

import (
	"fmt"
	"math/rand/v2"
	"os"

	"github.com/golang-jwt/jwt/v5"
)

// CalcToken generates a JWT token for a given username.
// It creates a new JWT token using the HS512 signing method and signs it
// with a secret derived from the environment variable "JWT_SECRET"
// concatenated with the username and a random integer.
// The function returns the signed token as a string and any error that may occur during the signing process.
// Note: Ensure that the "JWT_SECRET" environment variable is set before calling this function.
func CalcToken(username string) (string, error) {
	t := jwt.New(jwt.SigningMethodHS512)
	token, jerr := t.SignedString([]byte((os.Getenv("JWT_SECRET") + username + fmt.Sprint(rand.Int64N(10000000)))))
	return token, jerr
}
