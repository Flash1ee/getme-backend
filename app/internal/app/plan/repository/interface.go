package plan_repository

import (
	"getme-backend/internal/app/plan/entities"
	skill_entities "getme-backend/internal/app/skill/entities"
)

type Repository interface {
	//Create with Errors:
	// 		app.GeneralError with Errors
	// 			postgresql_utilits.DefaultErrDB
	Create(offerID int64, skills []skill_entities.Skill, plan entities.Plan) (*entities.Plan, error)
}
