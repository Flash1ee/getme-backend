package task_usecase

import (
	"database/sql"

	"getme-backend/internal/app"
	plan_repository "getme-backend/internal/app/plans/repository"
	"getme-backend/internal/app/task/dto"
	"getme-backend/internal/app/task/entities"
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

	res, err := u.taskRepository.Create(entities.Task{
		Name: sql.NullString{
			String: data.Name,
			Valid:  true,
		},
		Description: sql.NullString{
			String: data.Description,
			Valid:  true,
		},
		Deadline: sql.NullTime{
			Time:  data.Deadline,
			Valid: true,
		},
		PlanID: sql.NullInt64{
			Int64: data.PlanID,
			Valid: true,
		},
	})

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

	err = u.taskRepository.ApplyTask(entities.Task{
		ID: sql.NullInt64{
			Int64: data.ID,
		},
	})

	if err != nil {
		return err
	}

	return nil
}

////	GetPlansByRole with Errors:
////	postgresql_utilits.NotFound
//// 	plans_usecase.UnknownRole
//// 		app.GeneralError with Errors
//// 			postgresql_utilits.DefaultErrDB
//func (u *TaskUsecase) GetPlansByRole(userID int64, role string) ([]dto.PlansWithSkillsDTO, error) {
//	var plans []entities.PlanWithSkill
//	var err error
//	switch role {
//	case mentor:
//		plans, err = u.planRepository.GetByMentor(userID)
//	case mentee:
//		plans, err = u.planRepository.GetByMentee(userID)
//	default:
//		return nil, plans_usecase.UnknownRole
//	}
//	if err != nil {
//		return nil, err
//	}
//	res := filterPlansData(plans)
//
//	return dto.ToPlansWithSkillsUsecase(res), nil
//}
