package handler_factory

import (
	user_usecase "getme-backend/internal/app/user/usecase"
)

//go:generate mockgen -destination=mocks/mock_repository_factory.go -package=mock_repository_factory . RepositoryFactory

type UsecaseFactory interface {
	GetUserUsecase() user_usecase.Usecase
}
