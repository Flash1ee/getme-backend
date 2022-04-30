package usecase_factory

import repUser "getme-backend/internal/app/user/repository"

//go:generate mockgen -destination=mocks/mock_repository_factory.go -package=mock_repository_factory . RepositoryFactory

type RepositoryFactory interface {
	GetUserRepository() repUser.Repository
}
