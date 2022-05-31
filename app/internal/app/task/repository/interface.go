package task_repository

import "getme-backend/internal/app/task/entities"

type Repository interface {
	//Create with Errors:
	// 		app.GeneralError with Errors
	// 			postgresql_utilits.DefaultErrDB
	Create(task entities.Task) (int64, error)
}
