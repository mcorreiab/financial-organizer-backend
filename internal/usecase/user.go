package usecase

import (
	"fmt"
	"mcorreiab/financial-organizer-backend/internal/entities"
)

type UserUseCase struct {
	userRepository UserRepository
	SignInKey      string
}

type UserRepository interface {
	SaveUser(user entities.User) (string, error)
	FindUserByUsername(username string) (*entities.User, error)
}

type UserExistsError struct {
	Username string
}

func (u UserExistsError) Error() string {
	return fmt.Sprintf("User with username %s already exists", u.Username)
}

type InvalidCredentialsError struct{}

func (e InvalidCredentialsError) Error() string {
	return "Invalid username or password"
}

func NewUserUseCase(userRepository UserRepository, signInKey string) UserUseCase {
	return UserUseCase{userRepository, signInKey}
}

func (uc UserUseCase) SaveUser(username string, password string) (string, error) {
	err := uc.checkIfUserExists(username)

	if err != nil {
		return "", err
	}

	u, err := entities.NewUser(username, password)

	if err != nil {
		return "", err
	}

	id, err := uc.userRepository.SaveUser(u)
	if err != nil {
		return "", err
	}

	return id, nil
}

func (uc UserUseCase) checkIfUserExists(username string) error {
	user, err := uc.userRepository.FindUserByUsername(username)

	if user != nil {
		return UserExistsError{username}
	}

	return err
}

func (uc UserUseCase) GenerateLoginToken(username string, password string) (entities.Token, error) {
	user, err := uc.userRepository.FindUserByUsername(username)

	if err != nil {
		return entities.Token{}, err
	}

	if user == nil {
		return entities.Token{}, InvalidCredentialsError{}
	}

	isAuthenticated := user.CompareHashAndPassword(password)

	if !isAuthenticated {
		return entities.Token{}, InvalidCredentialsError{}
	}

	return entities.NewToken(uc.SignInKey, user.Id)
}
