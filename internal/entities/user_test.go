package entities

import (
	"testing"
)

type fakePasswordEncrypt struct {
	expectedPassword string
	expectedError    error
}

func (p fakePasswordEncrypt) Encrypt(password string) (string, error) {
	return p.expectedPassword, p.expectedError
}

type GenericError struct{}

func (e GenericError) Error() string {
	return "Generic error"
}

func TestCreateUserWithEncryptedPassword(t *testing.T) {
	input := make(map[string]string)
	input["password1"] = "expected password 1"
	input["password2"] = "expected password 2"

	for password, expectedPassword := range input {
		user, err := newUser("username", password, fakePasswordEncrypt{expectedPassword, nil})

		if err != nil {
			t.Error("Failed to encrypt password: ", err)
		}

		if user.Password != expectedPassword {
			t.Errorf("Encrypted password %s is different from expected %s", user.Password, expectedPassword)
		}
	}
}

func TestEncryptUserPasswordWithError(t *testing.T) {
	_, error := newUser("username", "password", fakePasswordEncrypt{"", GenericError{}})

	if error == nil {
		t.Error("Should throw an error")
	}
}
