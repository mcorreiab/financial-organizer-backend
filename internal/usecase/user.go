package usecase

import (
	"fmt"
	"mcorreiab/financial-organizer-backend/internal/entities"
)

type UserUseCase struct {
	userRepository UserRepository
	SignInKey      string
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

func (uc UserUseCase) SaveUser(username string, password string) (userId string, err error) {
	err = uc.checkIfUserExists(username)

	if err != nil {
		return
	}

	u, err := entities.NewUser(username, password)

	if err != nil {
		return
	}

	return uc.userRepository.SaveUser(u)
}

func (uc UserUseCase) checkIfUserExists(username string) error {
	if user, err := uc.userRepository.FindUserByUsername(username); user != nil {
		return UserExistsError{}
	} else {
		return err
	}
}

func (uc UserUseCase) GenerateLoginToken(username string, password string) (token entities.JwtToken, err error) {
	user, err := uc.userRepository.FindUserByUsername(username)

	if err != nil {
		return
	}

	if user == nil {
		err = InvalidCredentialsError{}
		return
	}

	if isAuthenticated := user.CompareHashAndPassword(password); !isAuthenticated {
		err = InvalidCredentialsError{}
		return
	}

	return entities.NewToken(uc.SignInKey).CreateJwt(user.Id)
}
