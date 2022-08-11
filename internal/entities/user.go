package entities

import (
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Username string
	Password string
}

type passwordEncrypter interface {
	Encrypt(password string) (string, error)
}

type bCryptPasswordEncrypt struct{}

func (p bCryptPasswordEncrypt) Encrypt(password string) (string, error) {
	pass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(pass), nil
}

func NewUser(username string, password string) (User, error) {
	return newUser(username, password, bCryptPasswordEncrypt{})
}

func newUser(username string, password string, encrypter passwordEncrypter) (User, error) {
	ep, err := encrypter.Encrypt(password)
	if err != nil {
		return User{}, err
	}

	return User{username, ep}, nil
}
