package main

import (
	"github.com/SnehilSundriyal/finances-manager/internal/models"
    repository "github.com/SnehilSundriyal/finances-manager/internal/repository/dbrepo"

    //    repository "github.com/SnehilSundriyal/finances-manager/internal/repository/dbrepo"
    "github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/render"
	"log"
    "net/http"

)

var	MyFinance models.PersonalFinance
var Expenses []models.Expense
const Income = 10000

func (app *application) GetFinances(context *gin.Context) {
	MyFinance.Income = Income
	err := render.WriteJSON(context.Writer, MyFinance)
	if err != nil {
		log.Println(err)
		return
	}
}

func (app *application) CreateNewExpense(context *gin.Context) {
	var expense models.Expense

	expense = repository.AddExpense(len(Expenses) + 1)
	MyFinance.TotalExpenses += expense.Amount
	MyFinance.Savings -= expense.Amount

	Expenses = append(Expenses, expense)

	context.JSON(http.StatusCreated, gin.H{
		"message": "new expense added succesfully",
		"error": nil,
	})
}



