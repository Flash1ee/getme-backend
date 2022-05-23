package skill_usecase

import "getme-backend/internal/app/skill/dto"

type Usecase interface {
	// GetAllSkills with Errors:
	//	app.GeneralError with Errors:
	//		postgresql_utilits.DefaultErrDB
	GetAllSkills() (*dto.SkillsUsecase, error)
}
