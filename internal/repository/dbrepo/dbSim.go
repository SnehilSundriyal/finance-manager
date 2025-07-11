package repository

import (
	"github.com/SnehilSundriyal/finances-manager/internal/models"
	"time"
)

func GetMyFinance() models.PersonalFinance {
	var myFinance models.PersonalFinance

	return myFinance
}

func AddExpense(ID int) models.Expense {
	var expense models.Expense

	expense.ID = ID
	expense.Name = "Pizzeria"
	expense.Type = models.ExpenseTypes[0]
	expense.Amount = 700
	expense.CreatedAt = time.Now()

	return expense
}
