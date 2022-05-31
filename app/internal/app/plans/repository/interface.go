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
	//GetByID with Errors:
	//		postgresql_utilits.NotFound
	// 		app.GeneralError with Errors
	// 			postgresql_utilits.DefaultErrDB
	GetByID(id int64) (*entities.Plan, error)
	//GetPlanByTaskID with Errors:
	// 		postgresql_utilits.NotFound
	// 		app.GeneralError with Errors
	// 			postgresql_utilits.DefaultErrDB
	GetPlanByTaskID(taskID int64) (*entities.Plan, error)
	//GetPlanWithMentorAndTasks with Errors:
	// 		app.GeneralError with Errors
	// 			postgresql_utilits.DefaultErrDB
	GetPlanWithMentorAndTasks(mentorID int64, taskID int64) ([]entities.PlanWithUserAndTask, error)
	//GetPlanWithMenteeAndTasks with Errors:
	// 		app.GeneralError with Errors
	// 			postgresql_utilits.DefaultErrDB
	GetPlanWithMenteeAndTasks(menteeID int64, taskID int64) ([]entities.PlanWithUserAndTask, error)
}
