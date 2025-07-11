package models

import (
    "time"
)

// PersonalFinance is the finance model
type PersonalFinance struct {
	ID 				int 		`json:"-"`
	Income 			int 		`json:"income"`
	Expenses 		[]Expense	`json:"expenses"`
	TotalExpenses 	int 		`json:"total_expenses"`
	Savings 		int			`json:"savings"`
	CreatedAt 		time.Time 	`json:"-"`
	UpdatedAt 		time.Time 	`json:"-"`
}

// Expense is the expenditure model
type Expense struct {
	ID 			int 		`json:"expense_id"`
	Name 		string      `json:"name"`
	Type 		string 		`json:"expense_type"`
	Amount 		int 		`json:"expense_amount"`
	CreatedAt 	time.Time 	`json:"created_at"`
}

var ExpenseTypes = []string{"Food", "Services", "Travel", "Splitwise", "Entertainment", "Miscellaneous"}

