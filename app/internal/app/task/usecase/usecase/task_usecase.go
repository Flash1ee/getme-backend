package task_usecase

import (
	"getme-backend/internal/app"
	plan_repository "getme-backend/internal/app/plans/repository"
	"getme-backend/internal/app/task/dto"
	task_repository "getme-backend/internal/app/task/repository"
	task_usecase "getme-backend/internal/app/task/usecase"
	"getme-backend/internal/pkg/usecase"
	postgresql_utilits "getme-backend/internal/pkg/utilits/postgresql"
)

type TaskUsecase struct {
	usecase.BaseUsecase
	taskRepository task_repository.Repository
	planRepository plan_repository.Repository
}

func NewTaskUsecase(repoTask task_repository.Repository, repoPlan plan_repository.Repository) *TaskUsecase {
	return &TaskUsecase{
		taskRepository: repoTask,
		planRepository: repoPlan,
	}
}

//Create with Errors:
//		task_usecase.UserHaveNotThisPlan
// 		app.GeneralError with Errors
// 			postgresql_utilits.DefaultErrDB
func (u *TaskUsecase) Create(mentorID int64, data dto.CreateTaskUsecasDTO) (int64, error) {
	plan, err := u.planRepository.GetByID(data.PlanID)
	if err != nil {
		if err == postgresql_utilits.NotFound {
			return app.InvalidInt, task_usecase.UserHaveNotThisPlan
		}
		return app.InvalidInt, err
	}

	if plan.MentorID != mentorID {
		return app.InvalidInt, task_usecase.UserHaveNotThisPlan
	}

	res, err := u.taskRepository.Create(*data.ToTasksEntities())

	if err != nil {
		return app.InvalidInt, err
	}

	return res, nil
}

//ApplyTask with Errors:
//		task_usecase.UserHaveNotThisTask
// 		app.GeneralError with Errors
// 			postgresql_utilits.DefaultErrDB
func (u *TaskUsecase) ApplyTask(mentorID int64, data dto.TaskUsecaseDTO) error {
	mentorId, err := u.taskRepository.GetMentorId(data.ID)
	if err != nil {
		if err == postgresql_utilits.NotFound {
			return task_usecase.UserHaveNotThisTask
		}
		return err
	}

	if mentorId != mentorID {
		return task_usecase.UserHaveNotThisTask
	}

	err = u.taskRepository.ApplyTask(*data.ToTasksEntities())

	if err != nil {
		return err
	}

	return nil
}
