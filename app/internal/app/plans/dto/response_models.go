package dto

//go:generate easyjson -all -disallow_unknown_fields response_models.go

//easyjson:json
type PlanResponse struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	About    string `json:"about"`
	MenteeID *int64 `json:"mentee_id,omitempty"`
	MentorID *int64 `json:"mentor_id,omitempty"`
}

//easyjson:json
type PlanResponseMentor struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	About    string `json:"about"`
	MenteeID *int64 `json:"mentee_id,omitempty"`
}

//easyjson:json
type PlanResponseMentee struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	About    string `json:"about"`
	MentorID *int64 `json:"mentor_id,omitempty"`
}

//easyjson:json
type PlanWithSkillsResponseMentor struct {
	PlanResponseMentor
	Skills []string `json:"skills"`
}

//easyjson:json
type PlanWithSkillsResponseMentee struct {
	PlanResponseMentee
	Skills []string `json:"skills"`
}

//easyjson:json
type PlansWithSkillsResponseMentor struct {
	Plans []PlanWithSkillsResponseMentor `json:"plans"`
}

//easyjson:json
type PlansWithSkillsResponseMentee struct {
	Plans []PlanWithSkillsResponseMentee `json:"plans"`
}

func ToPlanWithSkillResponseByMentor(data []PlansWithSkillsDTO) PlansWithSkillsResponseMentor {
	res := &PlansWithSkillsResponseMentor{
		Plans: make([]PlanWithSkillsResponseMentor, 0),
	}
	for _, val := range data {
		tmp := val.ToPlanWithSkillsResponseMentor()
		res.Plans = append(res.Plans, tmp)
	}
	return *res
}

func ToPlanWithSkillsResponseByMentee(data []PlansWithSkillsDTO) PlansWithSkillsResponseMentee {
	res := PlansWithSkillsResponseMentee{
		Plans: make([]PlanWithSkillsResponseMentee, 0),
	}
	for _, val := range data {
		tmp := val.ToPlanWithSkillsResponseMentee()
		res.Plans = append(res.Plans, tmp)
	}
	return res
}
