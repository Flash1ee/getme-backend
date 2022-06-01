package task_repository

import "getme-backend/internal/app/task/entities"

type Repository interface {
	//Create with Errors:
	// 		app.GeneralError with Errors
	// 			postgresql_utilits.DefaultErrDB
	Create(task entities.Task) (int64, error)

	//ApplyTask with Errors:
	// 		app.GeneralError with Errors
	// 			postgresql_utilits.DefaultErrDB
	ApplyTask(task entities.Task) error

	//GetMentorId with Errors:
	//		postgresql_utilits.NotFound
	// 		app.GeneralError with Errors
	// 			postgresql_utilits.DefaultErrDB
	GetMentorId(taskId int64) (int64, error)
}
