package main

import (
	"github.com/gin-gonic/gin"
)

func (app *application) routes() *gin.Engine {
	router := gin.Default()

	router.GET("/my-finances", app.GetFinances)

	router.GET("/expenses", app.GetExpenses)
	router.POST("/expenses", app.CreateNewExpense)

	router.GET("/expense", app.GetSingleExpense)
	router.PATCH("/expense", app.UpdateSingleExpense)
	router.DELETE("/expense", app.DeleteSingleExpense)

	return router
}