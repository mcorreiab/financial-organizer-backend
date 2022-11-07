package adapter

import (
	"database/sql"
	"mcorreiab/financial-organizer-backend/internal/entities"
)

type ExpenseRepository struct {
	db *sql.DB
}

func NewExpenseRepository(db *sql.DB) *ExpenseRepository {
	return &ExpenseRepository{db}
}

func (r *ExpenseRepository) SaveExpense(expense entities.Expense) (string, error) {
	return query{
		Table:      "expenses (expense_name, expense_value, expense_user_id)",
		Parameters: []any{expense.Name, expense.Value, expense.User.Id},
		db:         r.db,
	}.Insert()
}
