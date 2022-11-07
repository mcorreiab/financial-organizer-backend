package entities

import (
	"testing"

	"github.com/golang-jwt/jwt/v4"
	"github.com/stretchr/testify/require"
)

func TestDecodeTokenWithSuccess(t *testing.T) {
	mockUserId := "userId"
	token := NewToken("key")
	jwtToken, err := token.CreateJwt(mockUserId)

	if err != nil {
		panic(err)
	}

	userId := token.DecodeJwtToken(jwtToken.AccessToken)

	require.Equal(t, mockUserId, userId)
}

func TestDecodeTokenWithoutUserId(t *testing.T) {
	tk := NewToken("key")
	token, err := jwt.New(tk.signingMethod).SignedString([]byte("key"))

	if err != nil {
		panic(err)
	}

	userId := tk.DecodeJwtToken(token)

	require.Empty(t, userId)
}

func TestDecodeFakeToken(t *testing.T) {
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"

	userId := NewToken("key").DecodeJwtToken(token)

	require.Empty(t, userId)
}

func TestDecodeTokenInvalidFormat(t *testing.T) {
	token := "123"

	userId := NewToken("key").DecodeJwtToken(token)

	require.Empty(t, userId)
}
