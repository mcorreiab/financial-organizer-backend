package entities

import (
	"github.com/golang-jwt/jwt/v4"
)

func NewToken() (string, error) {
	return jwt.New(jwt.SigningMethodHS256).SignedString([]byte("abcd"))
}
