package entities

import (
	"testing"

	"github.com/golang-jwt/jwt/v4"
)

func TestDecodeTokenWithSuccess(t *testing.T) {
	mockUserId := "userId"
	token, err := NewToken("key", mockUserId)

	if err != nil {
		panic(err)
	}

	userId, err := DecodeToken("key", token.AccessToken)

	if err != nil {
		panic(err)
	}

	if userId != mockUserId {
		t.Errorf("Value %s is different from expected %s", userId, mockUserId)
	}
}

func TestDecodeTokenWithoutUserId(t *testing.T) {
	token, err := jwt.New(signingMethod).SignedString([]byte("key"))

	if err != nil {
		panic(err)
	}

	_, err = DecodeToken("key", token)

	if _, ok := err.(InvalidToken); !ok {
		t.Errorf("Error should be an invalid token. Received %s", err)
	}
}

func TestDecodeInvalidToken(t *testing.T) {
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"

	_, err := DecodeToken("key", token)

	if _, ok := err.(InvalidToken); !ok {
		t.Errorf("Error should be an invalid token. Received %s", err)
	}
}
