package entities

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type JwtToken struct {
	AccessToken string
	ExpiresIn   int
}

type Token struct {
	key           string
	signingMethod jwt.SigningMethod
}

func NewToken(key string) *Token {
	return &Token{key, jwt.SigningMethodHS256}
}

func (t Token) CreateJwt(userId string) (JwtToken, error) {
	claims := jwt.RegisteredClaims{
		Subject:   userId,
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
	}
	token, err := jwt.NewWithClaims(t.signingMethod, claims).SignedString([]byte(t.key))
	if err != nil {
		return JwtToken{}, err
	}

	return JwtToken{token, int((time.Duration(1) * time.Hour).Seconds())}, nil
}

func (tk Token) DecodeJwtToken(token string) string {
	parsedToken, _ := jwt.ParseWithClaims(token, &jwt.RegisteredClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(tk.key), nil
	})

	if claims, ok := parsedToken.Claims.(*jwt.RegisteredClaims); ok && parsedToken.Valid {
		if claims.Subject == "" {
			return ""
		}

		return claims.Subject
	}

	return ""
}
