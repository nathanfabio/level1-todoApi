package model

import "time"

type Task struct {
	ID        int       `db:"id" json:"id"`
	Title     string    `db:"title" json:"title"`
	Done      bool      `db:"done" json:"done"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
}

type User struct {
	ID int `db:"id" json:"id"`
	Email string `db:"email" json:"email"`
	Password string `db:"password" json:"password"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
}

// type LoginInput struct {
// 	Email    string `json:"email"`
// 	Password string `json:"password"`
// }