package usecase

import "mcorreiab/financial-organizer-backend/internal/entities"

type InvalidToken struct{}

func (t InvalidToken) Error() string {
	return "Token is invalid"
}

type AuthUsecase struct {
	signInKey      string
	userRepository UserRepository
}

func NewAuthUsecase(signKey string, userRepository UserRepository) *AuthUsecase {
	return &AuthUsecase{signKey, userRepository}
}

func (uc *AuthUsecase) ValidateToken(token string) (userId string, err error) {
	userId = entities.NewToken(uc.signInKey).DecodeJwtToken(token)

	if userId == "" {
		err = InvalidToken{}
		return
	}

	return userId, nil
}
