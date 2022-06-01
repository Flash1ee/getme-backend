package plans_usecase

import (
	"github.com/pkg/errors"

	"getme-backend/internal/app/plans/dto"
	"getme-backend/internal/app/plans/entities"
	plans_repository "getme-backend/internal/app/plans/repository"
	plans_usecase "getme-backend/internal/app/plans/usecase"
	"getme-backend/internal/pkg/usecase"
	postgresql_utilits "getme-backend/internal/pkg/utilits/postgresql"
)

const (
	mentor = "mentor"
	mentee = "mentee"
)

type PlanUsecase struct {
	usecase.BaseUsecase
	planRepository plans_repository.Repository
}

func NewPlanUsecase(repoPlan plans_repository.Repository) *PlanUsecase {
	return &PlanUsecase{
		planRepository: repoPlan,
	}
}

//	GetPlansByRole with Errors:
//	postgresql_utilits.NotFound
// 	plans_usecase.UnknownRole
// 		app.GeneralError with Errors
// 			postgresql_utilits.DefaultErrDB
func (u *PlanUsecase) GetPlansByRole(userID int64, role string) ([]dto.PlansWithSkillsDTO, error) {
	var plans []entities.PlanWithSkill
	var err error
	switch role {
	case mentor:
		plans, err = u.planRepository.GetByMentor(userID)
	case mentee:
		plans, err = u.planRepository.GetByMentee(userID)
	default:
		return nil, plans_usecase.UnknownRole
	}
	if err != nil {
		return nil, err
	}
	res := filterPlansData(plans)

	return dto.ToPlansWithSkillsUsecase(res), nil
}

//GetPlanWithTasks with Errors:
//	postgresql_utilits.NotFound
//	plans_usecase.PlanNotFound
//	plans_usecase.InvalidTaskID
// 		app.GeneralError with Errors
// 			postgresql_utilits.DefaultErrDB
func (u *PlanUsecase) GetPlanWithTasks(userID int64, planID int64) (dto.PlanWithTasksUsecaseDTO, error) {
	var err error
	var isMentor = true
	res := dto.PlanWithTasksUsecaseDTO{}

	entitiesRes := make([]entities.PlanWithUserAndTask, 0)
	plan, err := u.planRepository.GetPlanByTaskID(planID)
	if err != nil {
		if errors.Is(err, postgresql_utilits.NotFound) {
			return res, plans_usecase.PlanNotFound
		}
		return res, err
	}
	if userID == plan.MentorID {
		entitiesRes, err = u.planRepository.GetPlanWithMenteeAndTasks(userID, planID)
	} else if userID == plan.MenteeID {
		isMentor = false
		entitiesRes, err = u.planRepository.GetPlanWithMentorAndTasks(userID, planID)
	} else {
		return res, plans_usecase.InvalidTaskID
	}

	filtered := filterPlansByTasks(entitiesRes)
	res = dto.ToPlanWithTasksUsecaseDTO(filtered)[0]
	res.IsMentor = isMentor

	return res, nil
}
