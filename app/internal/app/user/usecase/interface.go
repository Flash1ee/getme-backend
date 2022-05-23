package user_usecase

import (
	"getme-backend/internal/app/user/dto"
	"getme-backend/internal/app/user/entities"
)

//go:generate mockgen -destination=mock/$GOFILE -package=mock -source=$GOFILE

type Usecase interface {
	FindByNickname(nickname string) (*entities.User, error)
	CreateBaseUser(nickname string) (int64, error)
	CreateFilledUser(data *dto.UserUsecase) (int64, error)
}
