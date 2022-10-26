package entities

import (
	"testing"

	"github.com/golang-jwt/jwt/v4"
)

func TestDecodeTokenWithSuccess(t *testing.T) {
	mockUserId := "userId"
	token := NewToken("key")
	jwtToken, err := token.CreateJwt(mockUserId)

	if err != nil {
		panic(err)
	}

	userId, err := token.DecodeJwtToken(jwtToken.AccessToken)

	if err != nil {
		panic(err)
	}

	if userId != mockUserId {
		t.Errorf("Value %s is different from expected %s", userId, mockUserId)
	}
}

func TestDecodeTokenWithoutUserId(t *testing.T) {
	tk := NewToken("key")
	token, err := jwt.New(tk.signingMethod).SignedString([]byte("key"))

	if err != nil {
		panic(err)
	}

	_, err = tk.DecodeJwtToken(token)

	if _, ok := err.(InvalidToken); !ok {
		t.Errorf("Error should be an invalid token. Received %s", err)
	}
}

func TestDecodeInvalidToken(t *testing.T) {
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"

	_, err := NewToken("key").DecodeJwtToken(token)

	if _, ok := err.(InvalidToken); !ok {
		t.Errorf("Error should be an invalid token. Received %s", err)
	}
}
