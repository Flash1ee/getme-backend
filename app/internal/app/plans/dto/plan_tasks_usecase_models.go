package dto

import "getme-backend/internal/app/task/dto"

type PlanWithTasksUsecaseDTO struct {
	PlansUsecaseDTO
	dto.TaskUsecaseDTO
}
