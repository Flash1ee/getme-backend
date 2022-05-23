package handler_factory

import (
	auth_usecase "getme-backend/internal/app/auth/usecase"
	token_usecase "getme-backend/internal/app/token/usecase"
	user_usecase "getme-backend/internal/app/user/usecase"
)

//go:generate mockgen -destination=mock/$GOFILE -package=mock -source=$GOFILE

type UsecaseFactory interface {
	GetUserUsecase() user_usecase.Usecase
	GetAuthUsecase() auth_usecase.Usecase
	GetTokenUsecase() token_usecase.Usecase
}
