package user_repository

import (
	"getme-backend/internal/app/user/entities"
)

//go:generate mockgen -destination=mock/$GOFILE -package=mock -source=$GOFILE

type Repository interface {
	FindByNickname(nickname string) (*entities.User, error)
	// CreateBaseUser Errors:
	// 		user_repository.EmailAlreadyExist
	// 		user_repository.NicknameAlreadyExist
	// 		app.GeneralError with Errors
	// 			postgresql_utilits.DefaultErrDB
	CreateBaseUser(nickname string) (int64, error)
	CreateFilledUser(data *entities.User) (int64, error)
}
