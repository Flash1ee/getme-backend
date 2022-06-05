package dto

import (
	"database/sql"
	"time"

	"getme-backend/internal/app/task/entities"
)

type CreateTaskUsecasDTO struct {
	Name        string
	Description string
	Deadline    time.Time
	PlanID      int64
}

func (model *CreateTaskUsecasDTO) ToTasksEntities() *entities.Task {
	return &entities.Task{
		Name: sql.NullString{
			String: model.Name,
			Valid:  true,
		},
		Description: sql.NullString{
			String: model.Description,
			Valid:  true,
		},
		Deadline: sql.NullTime{
			Time:  model.Deadline,
			Valid: true,
		},
		PlanID: sql.NullInt64{
			Int64: model.PlanID,
			Valid: true,
		},
	}
}

type TaskUsecaseDTO struct {
	ID          int64
	Name        string
	Description string
	Deadline    time.Time
	Status      string
}

func (model *TaskUsecaseDTO) ToTasksEntities() *entities.Task {
	return &entities.Task{
		ID: sql.NullInt64{
			Int64: model.ID,
			Valid: true,
		},
		Name: sql.NullString{
			String: model.Name,
			Valid:  true,
		},
		Description: sql.NullString{
			String: model.Description,
			Valid:  true,
		},
		Deadline: sql.NullTime{
			Time:  model.Deadline,
			Valid: true,
		},
	}
}
func (model *TaskUsecaseDTO) ToTasksResponse() *ResponseTask {
	return &ResponseTask{
		ID:          model.ID,
		Name:        model.Name,
		Description: model.Description,
		Deadline:    model.Deadline,
		Status:      model.Status,
	}
}
