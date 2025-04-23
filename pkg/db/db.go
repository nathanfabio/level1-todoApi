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

	taksSchema := `CREATE TABLE IF NOT EXISTS tasks (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
		done BOOLEAN NOT NULL DEFAULT false,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP);`

	userSchema := `CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		email TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP);`

	db.MustExec(taksSchema)
	db.MustExec(userSchema)
	return db
}