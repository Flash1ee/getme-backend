package dto

import "getme-backend/internal/app/user/dto"

//go:generate easyjson -all -disallow_unknown_fields request_models.go

//easyjson:json
type AuthTelegramCheckRequest struct {
	Token     string `query:"token" json:"token" validate:"required"`
	ID        int64  `query:"id" json:"id" validate:"required"`
	AuthDate  int64  `query:"auth_date" json:"auth_date" validate:"required"`
	FirstName string `query:"first_name" json:"first_name" validate:"required"`
	LastName  string `query:"last_name" json:"last_name" validate:"required"`
	Username  string `query:"username" json:"username" validate:"required"`
	Avatar    string `query:"photo_url" json:"photo_url"`
	Hash      string `query:"hash" json:"hash" validate:"required"`
}

func (req *AuthTelegramCheckRequest) ToAuthUsecase() *AuthUsecase {
	return &AuthUsecase{
		ID:        req.ID,
		AuthDate:  req.AuthDate,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Username:  req.Username,
		Avatar:    req.Avatar,
		Hash:      req.Hash,
	}
}
func (req *AuthTelegramCheckRequest) ToUserUsecase() *dto.UserUsecase {
	return &dto.UserUsecase{
		ID:        req.ID,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Nickname:  req.Username,
		Avatar:    req.Avatar,
	}
}

//easyjson:json
type AuthRequest struct {
	Token string `query:"token" json:"token" validate:"required"`
}

//easyjson:json
type AuthSimpleRequest struct {
	Login    string `json:"login" validate:"required"`
	Password string `json:"password" validate:"required"`
}

//easyjson:json
type AuthSimpleRegistrationRequest struct {
	Login    string `json:"login" validate:"required,alphanumunicode, min=5, max=25"`
	Password string `json:"password" validate:"required,min=4,max=50"`
}

func (req *AuthSimpleRegistrationRequest) ToSimpleRegistrationUsecase() *SimpleRegistrationUsecase {
	return &SimpleRegistrationUsecase{
		Login:    req.Login,
		Password: req.Password,
	}
}
