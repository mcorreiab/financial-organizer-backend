package entities

import (
	"github.com/golang-jwt/jwt/v4"
)

type Token struct {
	Key string
}

func (t Token) NewToken() (string, error) {
	return jwt.New(jwt.SigningMethodHS256).SignedString([]byte(t.Key))
}
