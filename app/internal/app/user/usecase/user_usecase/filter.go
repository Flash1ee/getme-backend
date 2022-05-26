package user_usecase

import "getme-backend/internal/app/user/entities"

func filterUsersData(users []entities.UserWithSkill) *entities.UserWithSkills {
	if len(users) == 0 {
		return nil
	}
	skills := make([]string, 0)
	res := &entities.UserWithSkills{}
	for _, val := range users {
		if val.Skill.Valid {
			skills = append(skills, val.Skill.String)
		}
	}
	res.User = users[0].User
	res.Skills = skills

	return res
}
