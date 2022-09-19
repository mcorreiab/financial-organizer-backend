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
	if pass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost); err != nil {
		return "", err
	} else {
		return string(pass), nil
	}
}

func NewUser(username, password string) (User, error) {
	return newUser(username, password, bCryptPasswordEncrypt{})
}

func newUser(username, password string, encrypter passwordEncrypter) (User, error) {
	if ep, err := encrypter.Encrypt(password); err != nil {
		return User{}, err
	} else {
		return User{username, ep}, nil
	}
}

func (u User) CompareHashAndPassword(password string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)); err == nil {
		return true
	}

	return false
}
