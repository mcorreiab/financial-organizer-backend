package entities

type Expense struct {
	Name  string
	Value int64
	User  User
}

func NewExpense(name string, value float64, user User) Expense {
	valueInCents := value * 100
	return Expense{name, int64(valueInCents), user}
}
