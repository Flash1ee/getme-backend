package repository_factory

import (
	"github.com/sirupsen/logrus"

	repUser "getme-backend/internal/app/user/repository"
	"getme-backend/internal/pkg/utilits"
)

type RepositoryFactory struct {
	expectedConnections utilits.ExpectedConnections
	logger              *logrus.Logger
	userRepository      repUser.Repository
}

func NewRepositoryFactory(logger *logrus.Logger, expectedConnections utilits.ExpectedConnections) *RepositoryFactory {
	return &RepositoryFactory{
		expectedConnections: expectedConnections,
		logger:              logger,
	}
}

func (f *RepositoryFactory) GetUserRepository() repUser.Repository {
	if f.userRepository == nil {
		//f.userRepository = repUserPsql.NewUserRepository(f.expectedConnections.SqlConnection)
	}
	return f.userRepository
}
