package entities

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type Token struct {
	Key    string
	UserId string
}

func (t Token) NewToken() (string, error) {
	claims := jwt.RegisteredClaims{
		Subject:   t.UserId,
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
	}
	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(t.Key))
}
