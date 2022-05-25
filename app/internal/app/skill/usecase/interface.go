package skill_usecase

import (
	"getme-backend/internal/app/skill/dto"
	dto2 "getme-backend/internal/app/user/dto"
)

type Usecase interface {
	// GetAllSkills with Errors:
	//	app.GeneralError with Errors:
	//		postgresql_utilits.DefaultErrDB
	GetAllSkills() (*dto.SkillsUsecase, error)
	// GetUsersBySkills with Errors:
	//	app.GeneralError with Errors:
	//		postgresql_utilits.DefaultErrDB
	GetUsersBySkills(data *dto.SkillsUsecase) ([]dto2.UserWithSkillsUsecase, error)
}
