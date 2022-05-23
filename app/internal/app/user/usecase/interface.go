package user_usecase

import (
	"getme-backend/internal/app/user/dto"
)

//go:generate mockgen -destination=mock/$GOFILE -package=mock -source=$GOFILE

type Usecase interface {
	// AuthTelegram Errors:
	//		ArgError
	//		BadAuth
	// 		app.GeneralError with Error
	// 			repository_postgresql.CreateError
	AuthTelegram(user *dto.UserAuthUsecase) (*dto.UserAuthUsecase, error)
	//AuthSimple Errors:
	//		IncorrectEmailOrPassword
	// 		postgresql_utilits.NotFound
	// 		app.GeneralError with Errors:
	// 			postgresql_utilits.DefaultErrDB
	AuthSimple(login string, password string) (int64, error)
	// CreateSimple Errors:
	//		models.EmptyPassword
	//		models.IncorrectNickname
	// 		models.IncorrectEmailOrPassword
	//		repository_postgresql.LoginAlreadyExist
	//		repository_postgresql.NicknameAlreadyExist
	//		UserExist
	// 		app.GeneralError with Errors
	// 			repository.DefaultErrDB
	CreateSimple(user *dto.UserSimpleRegistrationUsecase) (int64, error)
}
