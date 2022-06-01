package task_repository_postgresql

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"

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
	if err := repo.store.QueryRow(query, task.Name.String,
		task.Description.String, task.Deadline.Time, task.PlanID.Int64).Scan(&res); err != nil {
		return res, postgresql_utilits.NewDBError(err)
	}

	return res, nil
}

const updateTask = `
UPDATE task SET status = 'Выполнена' where id = ?`

//ApplyTask with Errors:
// 		app.GeneralError with Errors
// 			postgresql_utilits.DefaultErrDB
func (repo *TaskRepository) ApplyTask(task entities.Task) error {
	query := repo.store.Rebind(updateTask)

	if _, err := repo.store.Exec(query, task.ID); err != nil {
		return postgresql_utilits.NewDBError(err)
	}

	return nil
}

const getMentorId = `
SELECT mentor_id from plans
JOIN task t on plans.id = t.plan_id and t.id = ? LIMIT 1
`

//GetMentorId with Errors:
//		postgresql_utilits.NotFound
// 		app.GeneralError with Errors
// 			postgresql_utilits.DefaultErrDB
func (repo *TaskRepository) GetMentorId(taskId int64) (int64, error) {
	query := repo.store.Rebind(getMentorId)
	res := int64(-1)
	if err := repo.store.QueryRow(query, taskId).Scan(&res); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return res, postgresql_utilits.NotFound
		}
		return res, postgresql_utilits.NewDBError(err)
	}

	return res, nil
}
