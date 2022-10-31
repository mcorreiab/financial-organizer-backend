package usecase

import "mcorreiab/financial-organizer-backend/internal/entities"

type UserRepository interface {
	SaveUser(user entities.User) (string, error)
	FindUserByUsername(username string) (*entities.User, error)
	FindById(id string) (*entities.User, error)
}

type ExpenseRepository interface {
	SaveExpense(entities.Expense) (string, error)
}
