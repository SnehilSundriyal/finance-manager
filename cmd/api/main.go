package main

import (
	"context"
	"github.com/SnehilSundriyal/finances-manager/internal/repository"
	"github.com/SnehilSundriyal/finances-manager/internal/repository/dbrepo"
	"github.com/gin-gonic/gin"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
	"log"
	"os"
)

const port = ":8080"
type application struct {
	DB 		repository.DatabaseRepo
}

func main() {
	_ = godotenv.Load()
	var app application

	gin.SetMode(gin.ReleaseMode)

	// connecting to postgres database
	db, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL")) // DATABASE_URL="postgres://<YOUR_USERNAME_HERE>:<YOUR_PASSWORD_HERE>@localhost/postgres"
	if err != nil {
		log.Println("Failed to connect to database")
		log.Fatal(err)
	}
	app.DB = &dbrepo.PostgresDBRepo{DB: db}
	defer app.DB.Connect().Close(context.Background())

	log.Println("Connected to database....")


	// starting go server
	server := app.routes()

	log.Println("Starting GO application on port", port)

	err = server.Run(port)
	if err != nil {
		log.Fatal(err)
	}
}

