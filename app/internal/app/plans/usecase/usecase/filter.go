package plans_usecase

import (
	"getme-backend/internal/app/plans/entities"
	entities2 "getme-backend/internal/app/task/entities"
)

func filterPlansData(plans []entities.PlanWithSkill) []entities.PlanWithSkills {
	if len(plans) == 0 {
		return []entities.PlanWithSkills{}
	}
	planSkills := make(map[int64]*entities.PlanWithSkills)

	for _, val := range plans {
		if _, ok := planSkills[val.Plan.ID]; !ok {
			planSkills[val.Plan.ID] = &entities.PlanWithSkills{
				Plan:   val.Plan,
				Skills: []string{},
			}
		}
		if val.Skill.Valid {
			planSkills[val.Plan.ID].Skills = append(planSkills[val.Plan.ID].Skills, val.Skill.String)
		}
	}

	var res []entities.PlanWithSkills
	for _, plan := range planSkills {
		res = append(res, *plan)
	}

	return res
}

func filterPlansByTasks(plans []entities.PlanWithUserAndTask) []entities.PlanWithMentorAndTasks {
	if len(plans) == 0 {
		return []entities.PlanWithMentorAndTasks{}
	}

	planTasks := make(map[int64]*entities.PlanWithMentorAndTasks)
	for _, val := range plans {
		if _, ok := planTasks[val.Plan.ID]; !ok {
			planTasks[val.Plan.ID] = &entities.PlanWithMentorAndTasks{
				Plan:  val.Plan,
				User:  val.User,
				Tasks: make([]entities2.Task, 0),
			}
		}
		planTasks[val.Plan.ID].Tasks = append(planTasks[val.Plan.ID].Tasks, val.Task)
	}

	var res []entities.PlanWithMentorAndTasks
	for _, plan := range planTasks {
		res = append(res, *plan)
	}

	return res

}
