package entities

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type Token struct {
	AccessToken string
	ExpiresIn   int
}

type InvalidToken struct{}

func (t InvalidToken) Error() string {
	return "Token is invalid"
}

var signingMethod jwt.SigningMethod = jwt.SigningMethodHS256

func NewToken(key string, userId string) (Token, error) {
	claims := jwt.RegisteredClaims{
		Subject:   userId,
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
	}
	token, err := jwt.NewWithClaims(signingMethod, claims).SignedString([]byte(key))
	if err != nil {
		return Token{}, err
	}

	return Token{token, int((time.Duration(1) * time.Hour).Seconds())}, nil
}

func DecodeToken(key string, token string) (string, error) {
	parsedToken, _ := jwt.ParseWithClaims(token, &jwt.RegisteredClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	})

	if claims, ok := parsedToken.Claims.(*jwt.RegisteredClaims); ok && parsedToken.Valid {
		if claims.Subject == "" {
			return "", InvalidToken{}
		}

		return claims.Subject, nil
	}

	return "", InvalidToken{}
}
