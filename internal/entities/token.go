package entities

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type Token struct {
	AccessToken string
	ExpiresIn   int
}

func NewToken(key string, userId string) (Token, error) {
	claims := jwt.RegisteredClaims{
		Subject:   userId,
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
	}
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(key))
	if err != nil {
		return Token{}, err
	}

	return Token{token, int((time.Duration(1) * time.Hour).Seconds())}, nil
}
