package dto

//go:generate easyjson -all -disallow_unknown_fields response_models.go

//easyjson:json
type TaskIDResponse struct {
	ID int64 `json:"task_id"`
}
