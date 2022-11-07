package entities

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

type fakePasswordEncrypt struct {
	expectedPassword string
	expectedError    error
}

func (p fakePasswordEncrypt) Encrypt(password string) (string, error) {
	return p.expectedPassword, p.expectedError
}

const username = "username"

func TestCreateUserWithEncryptedPassword(t *testing.T) {
	input := make(map[string]string)
	input["password1"] = "expected password 1"
	input["password2"] = "expected password 2"

	for password, expectedPassword := range input {
		user, err := newUser(username, password, fakePasswordEncrypt{expectedPassword, nil})

		if err != nil {
			t.Error("Failed to encrypt password: ", err)
		}

		require.Equal(t, expectedPassword, user.Password)
	}
}

func TestEncryptUserPasswordWithError(t *testing.T) {
	_, err := newUser(username, "password", fakePasswordEncrypt{"", errors.New("")})

	require.NotNil(t, err)
}
