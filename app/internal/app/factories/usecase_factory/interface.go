package usecase_factory

import (
	auth_repository "getme-backend/internal/app/auth/repository"
	offer_repository "getme-backend/internal/app/offer/repository"
	skill_repository "getme-backend/internal/app/skill/repository"
	token_repository "getme-backend/internal/app/token/repository"
	repUser "getme-backend/internal/app/user/repository"
)

//go:generate mockgen -destination=mock/$GOFILE -package=mock -source=$GOFILE

type RepositoryFactory interface {
	GetUserRepository() repUser.Repository
	GetAuthRepository() auth_repository.Repository
	GetSkillRepository() skill_repository.Repository
	GetOfferRepository() offer_repository.Repository

	GetTokenRepository() token_repository.Repository
	GetTokenJWTRepository() token_repository.RepositoryJWT
}
