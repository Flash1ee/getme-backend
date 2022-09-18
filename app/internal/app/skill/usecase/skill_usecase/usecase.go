package skill_usecase

import (
	"fmt"

	"getme-backend/internal/app/skill/dto"
	skill_repository "getme-backend/internal/app/skill/repository"
	dto2 "getme-backend/internal/app/user/dto"
	"getme-backend/internal/app/user/entities"
	user_repository "getme-backend/internal/app/user/repository"
	"getme-backend/internal/pkg/usecase"

	"github.com/gomodule/redigo/redis"
	"github.com/mailru/easyjson"
)

type SkillUsecase struct {
	usecase.BaseUsecase
	log       logger
	skillRepo skill_repository.Repository
	usersRepo user_repository.Repository
}

func NewSkillUsecase(cache *redis.Pool, log logger, repo skill_repository.Repository, repoUser user_repository.Repository) *SkillUsecase {
	return &SkillUsecase{
		BaseUsecase: *usecase.NewBaseUsecase(cache),
		skillRepo:   repo,
		usersRepo:   repoUser,
		log:         log,
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
	skills := getSkillNameFromSkills(data)

	dataFromCache, err := u.StorageRedis.Get(skills)
	if err != nil {
		return nil, err
	}

	if dataFromCache != nil {
		res := make(dto2.UserWithSkillsUsecaseSlice, 2)
		err = easyjson.Unmarshal(dataFromCache, &res)
		if err != nil {
			return nil, fmt.Errorf("fail to unmarshal similar ids: %w", err)
		}
		u.log.Error("from cache")
		return res, nil
	}

	res, err := u.usersRepo.GetUsersBySkills(data.ToSkillEntites())
	if err != nil {
		return nil, err
	}
	resFinal := filterUsersData(res)

	respData := dto2.ToUsersWithSkillUsecase(resFinal)

	err = u.StorageRedis.Set(skills, respData)
	if err != nil {
		return respData, fmt.Errorf("can not set data to redis cache")
	}

	return respData, err
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

func getSkillNameFromSkills(data *dto.SkillsUsecase) string {
	var res string
	for _, val := range data.Skills {
		if res == "" {
			res = val.Name
			continue
		}
		res = fmt.Sprintf("%s,%s", res, val.Name)
	}
	return res
}
