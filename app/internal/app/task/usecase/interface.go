package task_usecase

import "getme-backend/internal/app/task/dto"

type Usecase interface {
	//Create with Errors:
	//		task_usecase.UserHaveNotThisPlan
	// 		app.GeneralError with Errors
	// 			postgresql_utilits.DefaultErrDB
	Create(mentorID int64, data dto.CreateTaskUsecasDTO) (int64, error)

	//ApplyTask with Errors:
	//		task_usecase.UserHaveNotThisTask
	// 		app.GeneralError with Errors
	// 			postgresql_utilits.DefaultErrDB
	ApplyTask(mentorID int64, data dto.TaskUsecaseDTO) error
}
