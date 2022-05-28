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
	// GetMentorsBySkills with Errors:
	//	app.GeneralError with Errors:
	//		postgresql_utilits.DefaultErrDB
	GetMentorsBySkills(data *dto.SkillsUsecase) ([]dto2.UserWithSkillsUsecase, error)
}
