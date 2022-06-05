package entities

import (
	"database/sql"
	"time"

	skill_entities "getme-backend/internal/app/skill/entities"
	entities_task "getme-backend/internal/app/task/entities"
	entities_user "getme-backend/internal/app/user/entities"
)

type Plan struct {
	ID        int64     `db:"id"`
	Name      string    `db:"name"`
	About     string    `db:"about"`
	IsActive  bool      `db:"is_active"`
	Progress  float64   `db:"progress"`
	MentorID  int64     `db:"mentor_id"`
	MenteeID  int64     `db:"mentee_id"`
	CreatedAt time.Time `db:"created_at"`
}

type PlansSkills struct {
	ID        int64  `db:"id"`
	PlanID    int64  `db:"plan_id"`
	SkillName string `db:"skill_name"`
}

type PlanWithSkill struct {
	Plan
	Skill sql.NullString `db:"skill_name"`
}

type PlanWithSkills struct {
	Plan
	Skills []string `db:"skill_name"`
}

func ToPlansSkills(planID int64, skills []skill_entities.Skill) []PlansSkills {
	res := make([]PlansSkills, 0, len(skills))
	for _, val := range skills {
		res = append(res, PlansSkills{
			PlanID:    planID,
			SkillName: val.Name,
		})
	}
	return res
}

type PlanWithUserAndTask struct {
	Plan
	entities_user.User
	entities_task.Task
}

type PlanWithMentorAndTasks struct {
	Plan
	entities_user.User
	Tasks []entities_task.Task
}
