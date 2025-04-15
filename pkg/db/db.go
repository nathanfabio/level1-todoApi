package db

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

func Connect() *sqlx.DB {
	db, err := sqlx.Open("sqlite3", "tasks.db")
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	

	schema := `CREATE TABLE IF NOT EXISTS tasks (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
		done BOOLEAN NOT NULL DEFAULT false,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP)`

	db.MustExec(schema)
	return db
}