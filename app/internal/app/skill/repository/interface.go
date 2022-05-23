package skill_repository

import "getme-backend/internal/app/skill/entities"

type Repository interface {
	// GetAllSkills Errors:
	// 		app.GeneralError with Errors:
	// 			repository.DefaultErrDB
	GetAllSkills() ([]entities.Skill, error)
}
