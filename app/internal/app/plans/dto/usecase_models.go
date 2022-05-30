package dto

import "getme-backend/internal/app/plans/entities"

type PlansCreateUsecaseDTO struct {
	ID       int64
	Name     string
	About    string
	MentorID int64
	MenteeID int64
}

func (m *PlansCreateUsecaseDTO) ToPlanResponseMentor() *PlanResponseMentor {
	return &PlanResponseMentor{
		ID:       m.ID,
		Name:     m.Name,
		About:    m.About,
		MenteeID: &m.MenteeID,
	}
}
func EntityToPlanCreateUsecaseDTO(plan *entities.Plan) *PlansCreateUsecaseDTO {
	return &PlansCreateUsecaseDTO{
		ID:       plan.ID,
		Name:     plan.Name,
		About:    plan.About,
		MentorID: plan.MentorID,
		MenteeID: plan.MenteeID,
	}
}

type PlansUsecaseDTO struct {
	ID       int64
	Name     string
	About    string
	Progress float64
	MentorID int64
	MenteeID int64
}

func (m *PlansUsecaseDTO) ToPlanResponse() *PlanResponse {
	return &PlanResponse{
		ID:       m.ID,
		Name:     m.Name,
		About:    m.About,
		MentorID: &m.MentorID,
		MenteeID: &m.MenteeID,
	}
}
func (m *PlansUsecaseDTO) ToPlanResponseMentor() *PlanResponseMentor {
	return &PlanResponseMentor{
		ID:       m.ID,
		Name:     m.Name,
		About:    m.About,
		Progress: m.Progress,
		MenteeID: &m.MenteeID,
	}
}
func (m *PlansUsecaseDTO) ToPlanResponseMentee() *PlanResponseMentee {
	return &PlanResponseMentee{
		ID:       m.ID,
		Name:     m.Name,
		About:    m.About,
		Progress: m.Progress,
		MentorID: &m.MentorID,
	}
}

type PlansWithSkillsDTO struct {
	PlansUsecaseDTO
	Skills []string
}

func (m *PlansWithSkillsDTO) ToPlanWithSkillsResponseMentor() PlanWithSkillsResponseMentor {
	return PlanWithSkillsResponseMentor{
		PlanResponseMentor: *m.PlansUsecaseDTO.ToPlanResponseMentor(),
		Skills:             m.Skills,
	}
}
func (m *PlansWithSkillsDTO) ToPlanWithSkillsResponseMentee() PlanWithSkillsResponseMentee {
	return PlanWithSkillsResponseMentee{
		PlanResponseMentee: *m.PlansUsecaseDTO.ToPlanResponseMentee(),
		Skills:             m.Skills,
	}
}

func ToPlansWithSkillsUsecase(data []entities.PlanWithSkills) []PlansWithSkillsDTO {
	res := make([]PlansWithSkillsDTO, 0)
	for _, val := range data {
		res = append(res, *ToPlansWithSkillUsecase(&val))
	}
	return res
}
func ToPlansWithSkillUsecase(data *entities.PlanWithSkills) *PlansWithSkillsDTO {
	return &PlansWithSkillsDTO{
		PlansUsecaseDTO: PlansUsecaseDTO{
			ID:       data.ID,
			Name:     data.Name,
			About:    data.About,
			Progress: data.Progress,
			MentorID: data.MentorID,
			MenteeID: data.MenteeID,
		},
		Skills: data.Skills,
	}
}
