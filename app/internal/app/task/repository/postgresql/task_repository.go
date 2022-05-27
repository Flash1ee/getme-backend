package task_postgresql

import (
	"github.com/jmoiron/sqlx"
)

type TaskRepository struct {
	store *sqlx.DB
}

func NewTaskRepository(store *sqlx.DB) *TaskRepository {
	return &TaskRepository{
		store: store,
	}
}
