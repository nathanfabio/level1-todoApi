package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/nathanfabio/level1-todoApi/internal/model"
)

type TaskRepository struct {
	db *sqlx.DB
}

func NewTaskRepository(db *sqlx.DB) *TaskRepository {
	return &TaskRepository{db: db}
}

func (r *TaskRepository) Create(task *model.Task) error {
	_, err := r.db.NamedExec(`INSERT INTO tasks (title, done) VALUES (:title, :done)`, task)
	return err
}

func (r *TaskRepository) GetAll() ([]model.Task, error) {
	var tasks []model.Task
	err := r.db.Select(&tasks, "SELECT * FROM tasks ORDER BY created_at DESC")
	return tasks, err
}

func (r *TaskRepository) UpdateStatus(id int, done bool) error {
	_, err := r.db.Exec(`UPDATE tasks SET done = ? WHERE id = ?`, done, id)
	return err
}