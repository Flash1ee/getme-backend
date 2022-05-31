package task_repository_postgresql

import (
	"github.com/jmoiron/sqlx"

	"getme-backend/internal/app/task/entities"
	postgresql_utilits "getme-backend/internal/pkg/utilits/postgresql"
)

type TaskRepository struct {
	store *sqlx.DB
}

func NewTaskRepository(store *sqlx.DB) *TaskRepository {
	return &TaskRepository{
		store: store,
	}
}

const queryCreateTask = `
INSERT INTO task(name, description, deadline, plan_id) VALUES (?, ?, ?, ?) RETURNING id`

//Create with Errors:
// 		app.GeneralError with Errors
// 			postgresql_utilits.DefaultErrDB
func (repo *TaskRepository) Create(task entities.Task) (int64, error) {
	query := repo.store.Rebind(queryCreateTask)
	res := int64(-1)
	if err := repo.store.QueryRow(query, task.Name, task.Description, task.Deadline, task.PlanID).Scan(&res); err != nil {
		return res, postgresql_utilits.NewDBError(err)
	}

	return res, nil
}
