package user_usecase

import (
	"getme-backend/internal/app/user/dto"
	"getme-backend/internal/app/user/entities"
)

//go:generate mockgen -destination=mock/$GOFILE -package=mock -source=$GOFILE

type Usecase interface {
	//	FindByID with Errors:
	// 		app.GeneralError with Errors
	// 			postgresql_utilits.DefaultErrDB
	FindByID(id int64) (*dto.UserWithSkillsUsecase, error)
	FindByNickname(nickname string) (*entities.User, error)
	CreateBaseUser(nickname string) (int64, error)
	CreateFilledUser(data *dto.UserUsecase) (int64, error)
	UpdateUser(user *dto.UserUsecase) (*dto.UserUsecase, error)
	//	GetMentorStatus with Errors:
	//		postgresql_utilits.NotFound
	//		app.GeneralError with Errors
	//			postgresql_utilits.DefaultErrDB
	GetMentorStatus(id int64) (*dto.UserStatusUsecase, error)
	//UpdateMentorStatus with Errors:
	//		postgresql_utilits.NotFound
	//		app.GeneralError with Errors
	//			postgresql_utilits.DefaultErrDB
	UpdateMentorStatus(user *dto.UserStatusUsecase) (*dto.UserStatusUsecase, error)
}
