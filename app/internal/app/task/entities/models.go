package entities

import "time"

type Task struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Deadline    time.Time `json:"deadline"`
	Status      string    `json:"status"`
	PlanID      int64     `json:"plan_id"`
	CreatedAt   time.Time `json:"created_at"`
}
