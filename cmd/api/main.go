package main

import (
	// "context"
	// "github.com/jackc/pgx/v5/pgxpool"
	// "github.com/joho/godotenv"
	"log"
	// "os"
)

const port = ":8080"
type application struct {
	DB string
}

func main() {
//	_ = godotenv.Load()
	var app application

	server := app.routes()

//	dbPool, err := pgxpool.New(context.Background(), os.Getenv("DATABASE_URL")) // DATABASE_URL="postgres://YOUR_USERNAME:YOUR_PASSWORD@localhost/postgres"
//	if err != nil {
//		log.Fatal("Failed to establish a connection to the database with error:", err)
//	}
//	defer dbPool.Close()

	log.Println("Starting GO application on port", port)

	err := server.Run(port)
	if err != nil {
		log.Fatal(err)
	}
}

