package db

import (
	"log"

	"github.com/jmoiron/sqlx"
)

func Connect() *sqlx.DB {
	db, err := sqlx.Open("sqlite3", "tasks.db")
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	defer db.Close()

	schema := `CREATE TABLE IF NOT EXISTS tasks (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
		done BOOLEAN NOT NULL DEFAULT false,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP)`

	db.MustExec(schema)
	return db
}