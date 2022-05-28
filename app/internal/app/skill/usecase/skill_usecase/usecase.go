package skill_usecase

import (
	"getme-backend/internal/app/skill/dto"
	skill_repository "getme-backend/internal/app/skill/repository"
	dto2 "getme-backend/internal/app/user/dto"
	"getme-backend/internal/app/user/entities"
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

// GetMentorsBySkills with Errors:
//	app.GeneralError with Errors:
//		postgresql_utilits.DefaultErrDB
func (u *SkillUsecase) GetMentorsBySkills(data *dto.SkillsUsecase) ([]dto2.UserWithSkillsUsecase, error) {
	res, err := u.usersRepo.GetUsersBySkills(data.ToSkillEntites())
	if err != nil {
		return nil, err
	}
	resFinal := filterUsersData(res)

	return dto2.ToUsersWithSkillUsecase(resFinal), nil
}

func filterUsersData(users []entities.UserWithSkill) []entities.UserWithSkills {
	ids := map[int64]struct{}{}
	skills := map[int64][]string{}
	res := make([]entities.User, 0)
	resFinal := make([]entities.UserWithSkills, 0)
	for _, val := range users {
		if val.Skill.Valid {
			skills[val.ID] = append(skills[val.ID], val.Skill.String)
		}
		if _, ok := ids[val.ID]; !ok {
			ids[val.ID] = struct{}{}
			res = append(res, val.User)
		}
	}
	for _, val := range res {
		resFinal = append(resFinal, entities.UserWithSkills{
			User:   val,
			Skills: skills[val.ID],
		})
	}
	return resFinal
}
