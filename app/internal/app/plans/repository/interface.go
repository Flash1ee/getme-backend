package plans_repository

import (
	"getme-backend/internal/app/plans/entities"
	skill_entities "getme-backend/internal/app/skill/entities"
)

type Repository interface {
	//Create with Errors:
	// 		app.GeneralError with Errors
	// 			postgresql_utilits.DefaultErrDB
	Create(offerID int64, skills []skill_entities.Skill, plan entities.Plan) (*entities.Plan, error)
	//GetByMentor with Errors:
	//		postgresql_utilits.NotFound
	// 		app.GeneralError with Errors
	// 			postgresql_utilits.DefaultErrDB
	GetByMentor(mentorID int64) ([]entities.PlanWithSkill, error)
	//GetByMentee with Errors:
	//		postgresql_utilits.NotFound
	// 		app.GeneralError with Errors
	// 			postgresql_utilits.DefaultErrDB
	GetByMentee(menteeID int64) ([]entities.PlanWithSkill, error)
}
