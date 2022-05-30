package handler_factory

import (
	auth_usecase "getme-backend/internal/app/auth/usecase"
	offer_usecase "getme-backend/internal/app/offer/usecase"
	plans_usecase "getme-backend/internal/app/plans/usecase"
	skill_usecase "getme-backend/internal/app/skill/usecase"
	token_usecase "getme-backend/internal/app/token/usecase"
	user_usecase "getme-backend/internal/app/user/usecase"
)

//go:generate mockgen -destination=mock/$GOFILE -package=mock -source=$GOFILE

type UsecaseFactory interface {
	GetUserUsecase() user_usecase.Usecase
	GetAuthUsecase() auth_usecase.Usecase
	GetTokenUsecase() token_usecase.Usecase
	GetSkillUsecase() skill_usecase.Usecase
	GetOfferUsecase() offer_usecase.Usecase
	GetPlansUsecase() plans_usecase.Usecase
}
