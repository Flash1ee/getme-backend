package dto

import "getme-backend/internal/app/user/entities"

type UserAuthUsecase struct {
	ID         int64  `json:"id"`
	TelegramID int64  `json:"tg_id"`
	AuthDate   int64  `json:"auth_date"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	Username   string `json:"username"`
	Avatar     string `json:"photo_url"`
	Hash       string `json:"hash"`
}

func (req *UserAuthUsecase) ToUserEntity() *entities.User {
	return &entities.User{
		TelegramID: req.TelegramID,
		FirstName:  req.FirstName,
		LastName:   req.LastName,
		Nickname:   req.Username,
		Avatar:     req.Avatar,
	}
}

func ToUserAuthUsecase(user *entities.User) *UserAuthUsecase {
	return &UserAuthUsecase{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Username:  user.Nickname,
		Avatar:    user.Avatar,
	}
}
