package dto

import "getme-backend/internal/app/user/dto"

//go:generate easyjson -all -disallow_unknown_fields response_models.go

//easyjson:json
type AuthResponse struct {
	ID       int64  `json:"id"`
	Nickname string `json:"nickname"`
}

//easyjson:json
type IDResponse struct {
	ID int64 `json:"id"`
}

//easyjson:json
type UpdateResponse struct {
	FirstName    string `json:"first_name,omitempty"`
	LastName     string `json:"last_name,omitempty"`
	Nickname     string `json:"nickname"`
	About        string `json:"about,omitempty"`
	Avatar       string `json:"avatar,omitempty"`
	IsSearchable bool   `json:"is_searchable"`
}

func ToUpdateResponseFromUsecase(usr *dto.UserUsecase) *UpdateResponse {
	return &UpdateResponse{
		FirstName:    usr.FirstName,
		LastName:     usr.LastName,
		Nickname:     usr.Nickname,
		Avatar:       usr.Avatar,
		About:        usr.About,
		IsSearchable: usr.IsSearchable,
	}
}
func ToUserResponseFromUsecase(data *AuthUsecase) *dto.UserResponse {
	return &dto.UserResponse{
		ID:        data.ID,
		FirstName: data.FirstName,
		LastName:  data.LastName,
		Avatar:    data.Avatar,
	}
}
