package entities

import "math/big"

type Expense struct {
	Name  string
	Value big.Float
	User  User
}
