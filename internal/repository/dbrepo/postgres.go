package repository

import (
	"database/sql"
	"time"
)

type PostgresDBRepo struct {
	DB *sql.DB
}

const dbTimeout = 3 * time.Second

func (db *PostgresDBRepo) Connect() *sql.DB {
	return db.DB
}