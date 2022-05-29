package entities

type Skill struct {
	Name  string `db:"name"`
	Color string `db:"color"`
}

func GetSkills(skills []string) []Skill {
	res := make([]Skill, 0, len(skills))
	for _, val := range skills {
		res = append(res, Skill{
			Name: val,
		})
	}
	return res
}
