package dto

import (
	"getme-backend/internal/app/plans/entities"
	"getme-backend/internal/app/task/dto"
	entities2 "getme-backend/internal/app/task/entities"
	dto2 "getme-backend/internal/app/user/dto"
)

type PlanWithTasksUsecaseDTO struct {
	PlansUsecaseDTO
	dto2.UserUsecase
	Tasks    []dto.TaskUsecaseDTO
	IsMentor bool
}

func ToTasksUsecaseDTO(tasks []entities2.Task) []dto.TaskUsecaseDTO {
	res := make([]dto.TaskUsecaseDTO, 0)
	for _, val := range tasks {
		res = append(res, ToTaskUsecaseDTO(val))
	}
	return res
}
func ToTaskUsecaseDTO(task entities2.Task) dto.TaskUsecaseDTO {
	if !task.ID.Valid {
		return dto.TaskUsecaseDTO{}
	}

	return dto.TaskUsecaseDTO{
		ID:          task.ID.Int64,
		Name:        task.Name.String,
		Description: task.Description.String,
		Deadline:    task.Deadline.Time,
		Status:      task.Status.String,
	}
}
func ToPlanWithTasksUsecaseDTO(plans []entities.PlanWithMentorAndTasks) []PlanWithTasksUsecaseDTO {
	res := make([]PlanWithTasksUsecaseDTO, 0, len(plans))

	for _, val := range plans {
		user := val.User
		tasks := val.Tasks
		plan := val.Plan
		planDTO := PlanWithTasksUsecaseDTO{
			PlansUsecaseDTO: PlansUsecaseDTO{
				ID:       plan.ID,
				Name:     plan.Name,
				About:    plan.About,
				Progress: plan.Progress,
				MentorID: plan.MentorID,
				MenteeID: plan.MenteeID,
			},
			UserUsecase: *dto2.ToUserUsecase(&user),
			Tasks:       ToTasksUsecaseDTO(tasks),
		}
		res = append(res, planDTO)
	}

	return res
}
