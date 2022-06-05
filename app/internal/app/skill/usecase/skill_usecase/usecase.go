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

func filterUsersData(users []entities_user.UserWithSkill) []entities_user.UserWithSkills {
	ids := map[int64]struct{}{}
	skills := map[int64][]string{}
	res := make([]entities_user.User, 0)
	resFinal := make([]entities_user.UserWithSkills, 0)
	for _, val := range users {
		if val.Skill.Valid {
			if val.User.ID.Valid {
				skills[val.ID.Int64] = append(skills[val.ID.Int64], val.Skill.String)
			}
		}
		if _, ok := ids[val.ID.Int64]; !ok {
			ids[val.ID.Int64] = struct{}{}
			res = append(res, val.User)
		}
	}
	for _, val := range res {
		resFinal = append(resFinal, entities_user.UserWithSkills{
			User:   val,
			Skills: skills[val.ID.Int64],
		})
	}
	return resFinal
}
