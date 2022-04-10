package repository_factory

import (
	"getme-backend/internal/app"
	repUser "getme-backend/internal/app/user/repository"
	repUserPsql "getme-backend/internal/app/user/repository/postgresql"

	"github.com/sirupsen/logrus"
)

type RepositoryFactory struct {
	expectedConnections app.ExpectedConnections
	logger              *logrus.Logger
	userRepository      repUser.Repository
}

func NewRepositoryFactory(logger *logrus.Logger, expectedConnections app.ExpectedConnections) *RepositoryFactory {
	return &RepositoryFactory{
		expectedConnections: expectedConnections,
		logger:              logger,
	}
}

func (f *RepositoryFactory) GetUserRepository() repUser.Repository {
	if f.userRepository == nil {
		f.userRepository = repUserPsql.NewUserRepository(f.expectedConnections.SqlConnection)
	}
	return f.userRepository
}
