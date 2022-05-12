package handler_factory

import (
	token_usecase "getme-backend/internal/app/token/usecase"
	user_usecase "getme-backend/internal/app/user/usecase"
)

//go:generate mockgen -destination=mock/$GOFILE -package=mock -source=$GOFILE

type UsecaseFactory interface {
	GetUserUsecase() user_usecase.Usecase
	GetTokenUsecase() token_usecase.Usecase
}
