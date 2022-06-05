package dto

import "time"

//go:generate easyjson -all -disallow_unknown_fields response_models.go

//easyjson:json
type TaskIDResponse struct {
	ID int64 `json:"task_id"`
}

//easyjson:json
type ResponseTask struct {
	ID          int64     `json:"id"`
	Name        string    `json:"title"`
	Description string    `json:"description"`
	Deadline    time.Time `json:"deadline"`
	Status      string    `json:"status"`
}
