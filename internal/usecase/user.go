package usecase

import "mcorreiab/financial-organizer-backend/internal/entities"

type SaveUserUseCase struct {
	userRepository UserRepository
}

type UserRepository interface {
	SaveUser(user entities.User) (string, error)
}

func NewSaveUserUseCase(userRepository UserRepository) SaveUserUseCase {
	return SaveUserUseCase{userRepository}
}

func (uc SaveUserUseCase) SaveUser(username string, password string) (string, error) {
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
