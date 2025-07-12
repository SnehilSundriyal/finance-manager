package repository

import (
	"github.com/SnehilSundriyal/finances-manager/internal/models"
	"github.com/jackc/pgx/v5"
)

type DatabaseRepo interface {
	Connect() *pgx.Conn
	GetMyFinance() (models.PersonalFinance, error)
	AddExpense(expense models.Expense) error
	GetExpenses() ([]models.Expense, error)
	UpdateFinances(myFinance models.PersonalFinance) error
}