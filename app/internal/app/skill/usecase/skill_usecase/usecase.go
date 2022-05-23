package skill_usecase

import (
	"getme-backend/internal/app/skill/dto"
	skill_repository "getme-backend/internal/app/skill/repository"
	"getme-backend/internal/pkg/usecase"
)

type SkillUsecase struct {
	usecase.BaseUsecase
	skillRepo skill_repository.Repository
}

func NewSkillUsecase(repo skill_repository.Repository) *SkillUsecase {
	return &SkillUsecase{
		skillRepo: repo,
	}
}

// GetAllSkills with Errors:
//	app.GeneralError with Errors:
//		postgresql_utilits.DefaultErrDB
func (u *SkillUsecase) GetAllSkills() (*dto.SkillsUsecase, error) {
	res, err := u.skillRepo.GetAllSkills()
	if err != nil {
		return nil, err
	}
	return dto.ToSkillUsecaseSlice(res), nil
}
