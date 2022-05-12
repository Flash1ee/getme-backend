package usecase_factory

import (
	token_repository "getme-backend/internal/app/token/repository"
	repUser "getme-backend/internal/app/user/repository"
)

//go:generate mockgen -destination=mock/$GOFILE -package=mock -source=$GOFILE

type RepositoryFactory interface {
	GetUserRepository() repUser.Repository
	GetTokenRepository() token_repository.Repository
	GetTokenJWTRepository() token_repository.RepositoryJWT
}
