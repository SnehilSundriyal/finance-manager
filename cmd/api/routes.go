package main

import (
	"github.com/gin-gonic/gin"
)

func (app *application) routes() *gin.Engine {
	router := gin.Default()

	router.GET("/", app.GetFinances)
	router.POST("/", app.CreateNewExpense)

	return router
}