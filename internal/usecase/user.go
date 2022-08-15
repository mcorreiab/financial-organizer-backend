package usecase

import (
	"fmt"
	"mcorreiab/financial-organizer-backend/internal/entities"
)

type SaveUserUseCase struct {
	userRepository UserRepository
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

func NewSaveUserUseCase(userRepository UserRepository) SaveUserUseCase {
	return SaveUserUseCase{userRepository}
}

func (uc SaveUserUseCase) SaveUser(username string, password string) (string, error) {
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

func (uc SaveUserUseCase) checkIfUserExists(username string) error {
	user, err := uc.userRepository.FindUserByUsername(username)

	if user != nil {
		return UserExistsError{username}
	}

	return err
}