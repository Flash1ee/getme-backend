package skill_usecase

import (
	"getme-backend/internal/app/skill/dto"
	skill_repository "getme-backend/internal/app/skill/repository"
	dto2 "getme-backend/internal/app/user/dto"
	user_repository "getme-backend/internal/app/user/repository"
	"getme-backend/internal/pkg/usecase"
)

type SkillUsecase struct {
	usecase.BaseUsecase
	skillRepo skill_repository.Repository
	usersRepo user_repository.Repository
}

func NewSkillUsecase(repo skill_repository.Repository, repoUser user_repository.Repository) *SkillUsecase {
	return &SkillUsecase{
		skillRepo: repo,
		usersRepo: repoUser,
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

// GetUsersBySkills with Errors:
//	app.GeneralError with Errors:
//		postgresql_utilits.DefaultErrDB
func (u *SkillUsecase) GetUsersBySkills(data *dto.SkillsUsecase) ([]dto2.UserUsecase, error) {
	res, err := u.usersRepo.GetUsersBySkills(data.ToSkillEntites())
	if err != nil {
		return nil, err
	}
	return dto2.ToUsersUsecase(res), nil
}
