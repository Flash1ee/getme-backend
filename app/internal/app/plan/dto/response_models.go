package dto

//go:generate easyjson -all -disallow_unknown_fields response_models.go

//easyjson:json
type ResponseAcceptPlan struct {
	ID       int64
	Name     string
	About    string
	MenteeID int64
	Skills   []string
}

func ToResponseAcceptPlan(plan *PlanCreateUsecaseDTO, skills []string) ResponseAcceptPlan {
	return ResponseAcceptPlan{
		ID:       plan.ID,
		Name:     plan.Name,
		About:    plan.About,
		MenteeID: plan.MenteeID,
		Skills:   skills,
	}
}
