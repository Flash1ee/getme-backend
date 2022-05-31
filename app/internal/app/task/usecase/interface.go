package task_usecase

import "getme-backend/internal/app/task/dto"

type Usecase interface {
	//Create with Errors:
	//		task_usecase.UserHaveNotThisPlan
	// 		app.GeneralError with Errors
	// 			postgresql_utilits.DefaultErrDB
	Create(mentorID int64, data dto.CreateTaskUsecasDTO) (int64, error)
}
