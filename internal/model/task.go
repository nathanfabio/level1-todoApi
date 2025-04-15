package model

import "time"

type Task struct {
	ID        int       `db:"id" json:"id"`
	Title     string    `db:"title" json:"title"`
	Done      bool      `db:"done" json:"done"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
}
