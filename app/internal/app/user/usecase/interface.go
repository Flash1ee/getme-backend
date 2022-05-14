package user_usecase

import (
	"getme-backend/internal/app/user/dto"
)

//go:generate mockgen -destination=mock/$GOFILE -package=mock -source=$GOFILE

type Usecase interface {
	// Auth Errors:
	//		ArgError
	//		BadAuth
	// 		app.GeneralError with Error
	// 			repository_postgresql.CreateError
	Auth(user *dto.UserAuthUsecase) (*dto.UserAuthUsecase, error)
}
