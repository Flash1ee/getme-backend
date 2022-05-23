package usecase_factory

import (
	auth_repository "getme-backend/internal/app/auth/repository"
	token_repository "getme-backend/internal/app/token/repository"
	repUser "getme-backend/internal/app/user/repository"
)

//go:generate mockgen -destination=mock/$GOFILE -package=mock -source=$GOFILE

type RepositoryFactory interface {
	GetUserRepository() repUser.Repository
	GetAuthRepository() auth_repository.Repository

	GetTokenRepository() token_repository.Repository
	GetTokenJWTRepository() token_repository.RepositoryJWT
}
