package util

import (
	"fmt"
	"math/rand/v2"
	"os"

	"github.com/golang-jwt/jwt/v5"
)

func CalcToken(username string) (string, error) {
	t := jwt.New(jwt.SigningMethodHS512)
	token, jerr := t.SignedString([]byte((os.Getenv("JWT_SECRET") + username + fmt.Sprint(rand.Int64N(10000000)))))
	return token, jerr
}
