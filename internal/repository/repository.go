package repository

import (
	"github.com/SnehilSundriyal/finances-manager/internal/models"
	"github.com/jackc/pgx/v5"
)

type DatabaseRepo interface {
	Connect() *pgx.Conn
	GetMyFinance() (models.PersonalFinance, error)
	GetExpenses() ([]models.Expense, error)
	AddExpense(expense models.Expense) error
	UpdateFinancesAfterExpense(myFinance models.PersonalFinance, addedAmount int,deletedAmount int) (models.PersonalFinance, error)
	UpdateExpense(expense models.Expense) (models.Expense, error)
	GetExpenseByID(id int) (models.Expense, error)
	DeleteExpense(id int) error
}