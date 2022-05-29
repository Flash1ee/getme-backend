package dto

import "getme-backend/internal/app/plan/entities"

type PlanCreateUsecaseDTO struct {
	ID       int64
	Name     string
	About    string
	MentorID int64
	MenteeID int64
}

func EntityToPlanCreateUsecaseDTO(plan *entities.Plan) *PlanCreateUsecaseDTO {
	return &PlanCreateUsecaseDTO{
		ID:       plan.ID,
		Name:     plan.Name,
		About:    plan.About,
		MentorID: plan.MentorID,
		MenteeID: plan.MenteeID,
	}
}
