package dto

import "getme-backend/internal/app/user/entities"

//go:generate easyjson -all -disallow_unknown_fields request_models.go

//easyjson:json
type UserUpdateRequest struct {
	Fullname string `json:"fulllname"`
	About    string `json:"about,omitempty"`
	Email    string `json:"email"`
}

func (req *UserUpdateRequest) ToUserEntities() *entities.User {
	return &entities.User{
		Fullname: req.Fullname,
		About:    req.About,
		Email:    req.Email,
	}
}
