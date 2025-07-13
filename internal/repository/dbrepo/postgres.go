package dbrepo

import (
	"context"
	"errors"
	"github.com/SnehilSundriyal/finances-manager/internal/models"
	"github.com/jackc/pgx/v5"
	"log"
	"time"
)

type PostgresDBRepo struct {
	DB *pgx.Conn
}

const dbTimeout = 3 * time.Second

func (db *PostgresDBRepo) Connect() *pgx.Conn {
	return db.DB
}

func (db *PostgresDBRepo) GetMyFinance() (models.PersonalFinance, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `
		SELECT ID, INCOME, TOTAL_EXPENSES, SAVINGS, CREATED_AT, UPDATED_AT
		FROM PERSONAL_FINANCE
`

	row := db.DB.QueryRow(ctx, query)

	var myFinance models.PersonalFinance
	err := row.Scan(
			&myFinance.ID,
			&myFinance.Income,
			&myFinance.TotalExpenses,
			&myFinance.Savings,
			&myFinance.CreatedAt,
			&myFinance.UpdatedAt,
		)

	if err != nil {
		log.Println(err)
		return models.PersonalFinance{}, err
	}

	return myFinance, nil
}

func (db *PostgresDBRepo) AddExpense(expense models.Expense) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `
		INSERT INTO EXPENSE (name, type, amount, created_at)
		VALUES ($1, $2, $3, $4)

`
	if expense.Name == "" {
		return errors.New("the expense name field cannot be empty")
	}

	if expense.Amount < 0 {
		return errors.New("the expense amount needs to be valid")
	}

	_, err := db.DB.Exec(ctx, query,
			&expense.Name,
			&expense.Type,
			&expense.Amount,
			&expense.CreatedAt,
		)

	if err != nil {
		return err
	}

	return nil
}

func (db *PostgresDBRepo) GetExpenses() ([]models.Expense, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `
		SELECT ID, NAME, TYPE, AMOUNT, CREATED_AT
		FROM EXPENSE
`

	rows, err := db.DB.Query(ctx, query)
	if err != nil {
		return []models.Expense{}, err
	}
	defer rows.Close()

	var expenses []models.Expense

	for rows.Next() {
		var expense models.Expense
		err := rows.Scan(
			&expense.ID,
			&expense.Name,
			&expense.Type,
			&expense.Amount,
			&expense.CreatedAt,
			)
		if err != nil {
			return []models.Expense{}, err
		}

		expenses = append(expenses, expense)
	}

	return expenses, nil
}

func (db *PostgresDBRepo) UpdateFinancesAfterExpense(myFinance models.PersonalFinance, addedAmount int,deletedAmount int) (models.PersonalFinance, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	myFinance.TotalExpenses += addedAmount
	myFinance.TotalExpenses -= deletedAmount
	myFinance.Savings = myFinance.Income - myFinance.TotalExpenses
	myFinance.UpdatedAt = time.Now()

	query := `
		UPDATE PERSONAL_FINANCE
		SET TOTAL_EXPENSES = $1, SAVINGS = $2, UPDATED_AT = $3
		WHERE ID = $4
`

	_, err := db.DB.Exec(ctx, query,
			myFinance.TotalExpenses,
			myFinance.Savings,
			myFinance.UpdatedAt,
			myFinance.ID,
		)

	if err != nil {
		return models.PersonalFinance{}, err
	}

	return myFinance, nil
}

func (db *PostgresDBRepo) GetExpenseByID(id int) (models.Expense, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()


	query := `
		SELECT ID, NAME, TYPE, AMOUNT, CREATED_AT
		FROM EXPENSE
		WHERE ID = $1
`

	row := db.DB.QueryRow(ctx, query, id)

	var expense models.Expense

	err := row.Scan(
			&expense.ID,
			&expense.Name,
			&expense.Type,
			&expense.Amount,
			&expense.CreatedAt,
		)
	if err != nil {
		return models.Expense{}, err
	}

	return expense, nil

}

func (db *PostgresDBRepo) UpdateExpense(expense models.Expense) (models.Expense, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `
		UPDATE EXPENSE
		SET NAME = $1, TYPE = $2, AMOUNT = $3
		WHERE ID = $4
`

	_, err := db.DB.Exec(ctx, query,
			expense.Name,
			expense.Type,
			expense.Amount,
			expense.ID,
		)
	if err != nil {
		return models.Expense{}, err
	}

	return expense, nil
}

func (db *PostgresDBRepo) DeleteExpense(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `
		DELETE FROM EXPENSE
		WHERE ID = $1
`

	_, err := db.DB.Exec(ctx, query, id)
	if err != nil {
		return err
	}

	return nil

}