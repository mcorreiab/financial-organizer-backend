package usecase

import (
	"mcorreiab/financial-organizer-backend/internal/entities"
)

type ExpenseUseCase struct {
	expenseRepository ExpenseRepository
	userRepository    UserRepository
	signInKey         string
}

func NewExpenseUseCase(expenseRepository ExpenseRepository, userRepository UserRepository, signKey string) *ExpenseUseCase {
	return &ExpenseUseCase{expenseRepository, userRepository, signKey}
}

func (uc *ExpenseUseCase) SaveExpense(name string, value float64, userId string) (expenseId string, err error) {
	user, err := uc.userRepository.FindById(userId)

	if err != nil {
		return
	}

	if user == nil {
		err = InvalidToken{}
		return
	}

	return uc.expenseRepository.SaveExpense(entities.NewExpense(name, value, *user))
}
