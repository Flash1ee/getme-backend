package dto

import (
	dto2 "getme-backend/internal/app/task/dto"
	"getme-backend/internal/app/user/dto"
)

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
	ID       int64   `json:"id"`
	Name     string  `json:"name"`
	About    string  `json:"about"`
	Progress float64 `json:"progress"`
	MenteeID *int64  `json:"mentee_id,omitempty"`
}

//easyjson:json
type PlanResponseMentee struct {
	ID       int64   `json:"id"`
	Name     string  `json:"name"`
	About    string  `json:"about"`
	Progress float64 `json:"progress"`
	MentorID *int64  `json:"mentor_id,omitempty"`
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

//easyjson:json
type PlanWithTaskResponseMentor struct {
	Title            string  `json:"title"`
	Description      string  `json:"description"`
	Progress         float64 `json:"progress"`
	dto.UserResponse `json:"mentor"`
	Tasks            []dto2.ResponseTask `json:"tasks,omitempty"`
}

//easyjson:json
type PlanWithTaskResponseMentee struct {
	Title            string  `json:"title"`
	Description      string  `json:"description"`
	Progress         float64 `json:"progress"`
	dto.UserResponse `json:"mentee"`
	Tasks            []dto2.ResponseTask `json:"tasks,omitempty"`
}

func ToPlanWithTaskResponseMentor(data PlanWithTasksUsecaseDTO) PlanWithTaskResponseMentor {
	res := PlanWithTaskResponseMentor{
		Title:        data.PlansUsecaseDTO.Name,
		Description:  data.PlansUsecaseDTO.About,
		Progress:     data.PlansUsecaseDTO.Progress,
		UserResponse: dto.ToUserResponse(data.UserUsecase),
		Tasks:        make([]dto2.ResponseTask, 0),
	}
	for _, val := range data.Tasks {
		res.Tasks = append(res.Tasks, *val.ToTasksResponse())
	}
	return res
}
func ToPlanWithTaskResponseMentee(data PlanWithTasksUsecaseDTO) PlanWithTaskResponseMentee {
	res := PlanWithTaskResponseMentee{
		Title:        data.PlansUsecaseDTO.Name,
		Description:  data.PlansUsecaseDTO.About,
		Progress:     data.PlansUsecaseDTO.Progress,
		UserResponse: dto.ToUserResponse(data.UserUsecase),
		Tasks:        make([]dto2.ResponseTask, 0),
	}
	for _, val := range data.Tasks {
		res.Tasks = append(res.Tasks, *val.ToTasksResponse())
	}
	return res
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
