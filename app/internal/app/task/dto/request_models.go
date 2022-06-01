package dto

import "time"

//go:generate easyjson -all -disallow_unknown_fields request_models.go

//easyjson:json
type RequestCreateTask struct {
	Name        string    `json:"name" validate:"required"`
	Description string    `json:"description" validate:"required"`
	Deadline    time.Time `json:"deadline" validate:"required,datetime"`
	PlanID      int64     `json:"plan_id,omitempty" validate:"required,gte=1"`
}

func (req *RequestCreateTask) ToRequestCreateTaskDTO() *CreateTaskUsecasDTO {
	return &CreateTaskUsecasDTO{
		Name:        req.Name,
		Description: req.Description,
		Deadline:    req.Deadline,
		PlanID:      req.PlanID,
	}
}
