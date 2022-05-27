package skill_repository

import "getme-backend/internal/app/skill/entities"

type Repository interface {
	// GetAllSkills Errors:
	// 		app.GeneralError with Errors:
	// 			repository.DefaultErrDB
	GetAllSkills() ([]entities.Skill, error)
	// CheckExists Errors:
	//		postgresql_utilits.Exists
	//		postgresql_utilits.NotFound
	// 		app.GeneralError with Errors:
	// 			postgresql_utilits.DefaultErrDB
	CheckExists(skillName string) error
}
