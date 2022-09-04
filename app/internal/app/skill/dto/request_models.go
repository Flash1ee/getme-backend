package dto

//go:generate easyjson -all -disallow_unknown_fields request_models.go

//easyjson:json
type RequestUsersBySkills struct {
	Skills []string `query:"skills"`
}

func (req *RequestUsersBySkills) ToSkillUsecase() *SkillsUsecase {
	res := &SkillsUsecase{
		Skills: make([]SkillUsecase, 0, len(req.Skills)),
	}
	for _, val := range req.Skills {
		res.Skills = append(res.Skills, SkillUsecase{
			Name: val,
		})
	}
	return res
}
