package plans_usecase

import "getme-backend/internal/app/plans/entities"

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
