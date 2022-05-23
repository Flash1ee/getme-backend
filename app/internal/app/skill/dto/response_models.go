package dto

//go:generate easyjson -all -disallow_unknown_fields response_models.go

//json:easyjson
type SkillResponse struct {
	Name  string `json:"name"`
	Color string `json:"color,omitempty"`
}

//json:easyjson
type SkillsResponse struct {
	Skills []SkillResponse `json:"skills"`
}

func ToSkillsResponseFromUsecase(usr *SkillsUsecase) SkillsResponse {
	res := SkillsResponse{
		Skills: make([]SkillResponse, 0, len(usr.Skills)),
	}
	for _, val := range usr.Skills {
		res.Skills = append(res.Skills, SkillResponse{
			Name:  val.Name,
			Color: val.Color,
		})
	}
	return res
}
