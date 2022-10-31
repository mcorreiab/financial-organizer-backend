package usecase

import (
	"math/big"
	"mcorreiab/financial-organizer-backend/internal/entities"
)

type InvalidToken struct{}

func (t InvalidToken) Error() string {
	return "Token is invalid"
}

type ExpenseUseCase struct {
	ExpenseRepository ExpenseRepository
	UserRepository    UserRepository
	SignInKey         string
}

func NewExpenseUseCase(expenseRepository ExpenseRepository, userRepository UserRepository, signKey string) *ExpenseUseCase {
	return &ExpenseUseCase{expenseRepository, userRepository, signKey}
}

func (uc *ExpenseUseCase) SaveExpense(name string, value big.Float, token string) (expenseId string, err error) {
	userId := entities.NewToken(uc.SignInKey).DecodeJwtToken(token)

	if userId == "" {
		err = InvalidToken{}
		return
	}

	user, err := uc.UserRepository.FindById(userId)

	if err != nil {
		return
	}

	if user == nil {
		err = InvalidToken{}
		return
	}

	return uc.ExpenseRepository.SaveExpense(entities.Expense{Name: name, Value: value, User: *user})
}
