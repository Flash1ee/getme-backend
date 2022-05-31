package plans_usecase

import (
	"getme-backend/internal/app/plans/dto"
	"getme-backend/internal/app/plans/entities"
	plans_repository "getme-backend/internal/app/plans/repository"
	plans_usecase "getme-backend/internal/app/plans/usecase"
	"getme-backend/internal/pkg/usecase"
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

func (u *PlanUsecase) GetPlanWithTasks(userID int64, taskID int64) (*dto.PlanWithTasksUsecaseDTO, error) {
	//var err error

	/*entitiesRes := make([]entities.PlanWithUserAndTask, 0)
	plan, err := u.planRepository.GetPlanByTaskID(taskID)
	if err != nil {
		if errors.Is(err, postgresql_utilits.NotFound) {
			return nil, plans_usecase.PlanNotFound
		}
		return nil, err
	}
	if userID == plan.MentorID {
		entitiesRes, err = u.planRepository.GetPlanWithMentorAndTasks(userID, taskID)
	} else if userID == plan.MenteeID {
		entitiesRes, err = u.planRepository.GetPlanWithMenteeAndTasks(userID, taskID)
	} else {
		return nil, plans_usecase.InvalidTaskID
	}
	*/
	//res := filterPlansByTasks(entitiesRes)
	// @TODO конвертнуть в []dto.PlanWithTasksUsecaseDTO
	return nil, nil
}
