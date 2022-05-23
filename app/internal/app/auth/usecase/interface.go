package auth_usecase

import (
	"getme-backend/internal/app/auth/dto"
)

//go:generate mockgen -destination=mock/$GOFILE -package=mock -source=$GOFILE

type Usecase interface {
	//AuthSimple Errors:
	//		auth_usecase.IncorrectLoginOrPassword
	// 		postgresql_utilits.NotFound
	// 		app.GeneralError with Errors:
	// 			postgresql_utilits.DefaultErrDB
	AuthSimple(login string, password string) (int64, error)

	//AuthTelegram Errors:
	//		ArgError
	//		BadAuth
	//		app.GeneralError with Error
	//			repository_postgresql.CreateError
	AuthTelegram(user *dto.AuthUsecase) (*dto.AuthUsecase, error)

	// CreateSimple Errors:
	//		models.EmptyPassword
	//		models.IncorrectNickname
	// 		models.IncorrectEmailOrPassword
	//		repository_postgresql.LoginAlreadyExist
	//		repository_postgresql.NicknameAlreadyExist
	//		UserExist
	// 		app.GeneralError with Errors
	// 			repository.DefaultErrDB
	CreateSimple(user *dto.SimpleRegistrationUsecase) (int64, error)
	//CreateTelegram(user *dto.UserSimpleRegistrationUsecase) (int64, error)
}
