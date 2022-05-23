package dto

import "time"

//go:generate easyjson -all -disallow_unknown_fields response_models.go

//easyjson:json
type UserResponse struct {
	FirstName    string    `json:"first_name,omitempty"`
	LastName     string    `json:"last_name,omitempty"`
	Nickname     string    `json:"nickname"`
	About        string    `json:"about,omitempty"`
	Avatar       string    `json:"avatar,omitempty"`
	IsSearchable bool      `json:"is_mentor"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

func ToUserResponse(user *UserUsecase) UserResponse {
	return UserResponse{
		FirstName:    user.FirstName,
		LastName:     user.LastName,
		Nickname:     user.Nickname,
		About:        user.About,
		Avatar:       user.Avatar,
		IsSearchable: user.IsSearchable,
		CreatedAt:    user.CreatedAt,
		UpdatedAt:    user.UpdatedAt,
	}
}
