package dto

import "getme-backend/internal/app/skill/entities"

type SkillUsecase struct {
	Name  string
	Color string
}

type SkillsUsecase struct {
	Skills []SkillUsecase `json:"skills"`
}

func ToSkillUsecase(data *entities.Skill) *SkillUsecase {
	return &SkillUsecase{
		Name:  data.Name,
		Color: data.Color,
	}
}
func ToSkillUsecaseSlice(data []entities.Skill) *SkillsUsecase {
	res := &SkillsUsecase{
		Skills: make([]SkillUsecase, 0, len(data)),
	}

	for _, val := range data {
		res.Skills = append(res.Skills, SkillUsecase{
			Name:  val.Name,
			Color: val.Color,
		})
	}
	return res
}
