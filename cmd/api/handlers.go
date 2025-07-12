package main

import (
	"github.com/SnehilSundriyal/finances-manager/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/render"
	"log"
    "net/http"
)

func (app *application) GetFinances(context *gin.Context) {
	MyFinance, err := app.DB.GetMyFinance()
	if err != nil {
		context.JSON(http.StatusNoContent, gin.H{
			"message": "error getting finance",
			"error": err,
		})

		return
	}

	err = render.WriteJSON(context.Writer, &MyFinance)
	if err != nil {
		log.Println(err)
		return
	}
}

func (app *application) CreateNewExpense(context *gin.Context) {
	var expense models.Expense
	err := context.ShouldBindJSON(&expense)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Error binding request as json",
			"error": err.Error(),
		})
		return
	}

	err = app.DB.AddExpense(expense)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": err,
			"error": err.Error(),
		})
		return
	}

	MyFinance, err := app.DB.GetMyFinance()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to get current finances",
			"error": err.Error(),
		})
		return
	}

	MyFinance, err = app.DB.UpdateFinancesAfterExpense(MyFinance, expense.Amount, 0)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to update finances",
			"error": err.Error(),
		})
		return
	}

	context.JSON(http.StatusCreated, gin.H{
		"message": "new expense added succesfully & finances updated",
		"new_expense_name": expense.Name,
		"new_expense_type": expense.Type,
		"new_expense_amount": expense.Amount,
	})
}

func (app *application) GetExpenses(context *gin.Context) {
	var Expenses []models.Expense
	var err error

	Expenses, err = app.DB.GetExpenses()
	if err != nil {
		context.JSON(http.StatusNoContent, gin.H{
			"message": "error getting expenses",
			"error": err,
		})

		return
	}

	err = render.WriteJSON(context.Writer, &Expenses)
	if err != nil {
		log.Println(err)
		return
	}

}

func (app *application) GetSingleExpense(context *gin.Context) {
	var expense models.Expense
	err := context.ShouldBindJSON(&expense)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "error binding json",
			"error": err.Error(),
		})
		return
	}

	expense, err = app.DB.GetExpenseByID(expense.ID)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "error getting expense",
			"error": err.Error(),
		})
		return
	}

	err = render.WriteJSON(context.Writer, &expense)
	if err != nil {
		log.Println(err)
		return
	}
}

func (app *application) UpdateSingleExpense(context *gin.Context) {
	var expense models.Expense
	err := context.ShouldBindJSON(&expense)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "error binding json",
			"error": err.Error(),
		})
		return
	}

	originalExpense, err := app.DB.GetExpenseByID(expense.ID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to get original expense",
			"error": err.Error(),
			})
		return
	}

	MyFinance, err := app.DB.GetMyFinance()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to get current finances",
			"error": err.Error(),
			})
		return
	}

	expense, err = app.DB.UpdateExpense(expense)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "error updating expense",
			"error": err.Error(),
		})
		return
	}

	MyFinance, err = app.DB.UpdateFinancesAfterExpense(MyFinance, expense.Amount, originalExpense.Amount)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to get update finances",
			"error": err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "expense updated succesfully & finances updated",
		"new_expense_name": expense.Name,
		"new_expense_type": expense.Type,
		"new_expense_amount": expense.Amount,
	})
}


