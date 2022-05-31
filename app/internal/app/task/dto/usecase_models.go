package dto

import "time"

type CreateTaskUsecasDTO struct {
	Name        string
	Description string
	Deadline    time.Time
	PlanID      int64
}

type TaskUsecaseDTO struct {
	ID          int64
	Name        string
	Description string
	Deadline    time.Time
	Status      string
}
