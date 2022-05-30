package plans_usecase

import (
	"getme-backend/internal/app/plans/dto"
)

type Usecase interface {
	//	GetPlansByRole with Errors:
	//	postgresql_utilits.NotFound
	// 	plans_usecase.UnknownRole
	// 		app.GeneralError with Errors
	// 			postgresql_utilits.DefaultErrDB
	GetPlansByRole(userID int64, role string) ([]dto.PlansWithSkillsDTO, error)
}
